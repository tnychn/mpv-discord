package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/tnychn/mpv-discord/discordrpc"
	"github.com/tnychn/mpv-discord/mpvrpc"
)

var client *mpvrpc.Client
var presence *discordrpc.Presence

func init() {
	log.SetFlags(log.Ltime | log.Lmsgprefix)
	log.SetPrefix("[discord] ")

	client = mpvrpc.NewClient(os.Args[1])
	presence = discordrpc.NewPresence("737663962677510245")
}

func getActivity() (activity discordrpc.Activity, err error) {
	getProperty := func(key string) (prop interface{}) {
		prop, err = client.GetProperty(key)
		return
	}

	getPropertyString := func(key string) (prop string) {
		prop, err = client.GetPropertyString(key)
		return
	}

	// Large Image
	activity.LargeImageKey = "mpv"
	activity.LargeImageText = "mpv"
	if version := getPropertyString("mpv-version"); version != "" {
		activity.LargeImageText += " " + version[4:]
	}

	// Details
	activity.Details = getPropertyString("media-title")
	metaTitle := getProperty("metadata/by-key/Title")
	metaArtist := getProperty("metadata/by-key/Artist")
	metaAlbum := getProperty("metadata/by-key/Album")
	if metaTitle != nil {
		activity.Details = metaTitle.(string)
	}

	// State
	if metaArtist != nil {
		activity.State += " by " + metaArtist.(string)
	}
	if metaAlbum != nil {
		activity.State += " on " + metaAlbum.(string)
	}
	if activity.State == "" {
		if aid, ok := getProperty("aid").(string); !ok || aid != "no" {
			activity.State += "A"
		}
		if vid, ok := getProperty("vid").(string); !ok || vid != "no" {
			activity.State += "V"
		}
		_timePos := getProperty("time-pos")
		_duration := getProperty("duration")
		if _timePos != nil && _duration != nil {
			timePos := time.Unix(int64(_timePos.(float64)), 0).UTC().Format("15:04:05")
			duration := time.Unix(int64(_duration.(float64)), 0).UTC().Format("15:04:05")
			activity.State += fmt.Sprintf(": %s/%s", timePos, duration)
		}
	}

	// Small Image
	buffering := getProperty("paused-for-cache")
	pausing := getProperty("pause")
	loopingFile := getPropertyString("loop-file")
	loopingPlaylist := getPropertyString("loop-playlist")
	if buffering != nil && buffering.(bool) {
		activity.SmallImageKey = "buffer"
		activity.SmallImageText = "Buffering"
	} else if pausing != nil && pausing.(bool) {
		activity.SmallImageKey = "pause"
		activity.SmallImageText = "Paused"
	} else if loopingFile != "no" || loopingPlaylist != "no" {
		activity.SmallImageKey = "loop"
		activity.SmallImageText = "Looping"
	} else {
		activity.SmallImageKey = "play"
		activity.SmallImageText = "Playing"
	}
	if percentage := getProperty("percent-pos"); percentage != nil {
		activity.SmallImageText += fmt.Sprintf(" (%d%%)", int(percentage.(float64)))
	}
	if pcount := getProperty("playlist-count"); pcount != nil && int(pcount.(float64)) > 1 {
		if ppos := getProperty("playlist-pos-1"); ppos != nil {
			activity.SmallImageText += fmt.Sprintf(" [%d/%d]", int(ppos.(float64)), int(pcount.(float64)))
		}
	}

	// Timestamps
	if timeRemaining := getPropertyString("time-remaining"); timeRemaining != "" {
		d, _ := time.ParseDuration(timeRemaining + "s")
		endTime := time.Now().Add(d)
		activity.Timestamps = &discordrpc.ActivityTimestamps{End: &endTime}
	}
	return
}

func main() {
	if err := presence.Open(); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := presence.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := client.Open(); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for range time.Tick(time.Second) {
		activity, err := getActivity()
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				break
			}
			log.Println(err)
			continue
		}
		if err = presence.Update(activity); err != nil {
			log.Println(err)
			continue
		}
	}
}
