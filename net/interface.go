package net

import "context"

type VideoNetwork interface {
	NewBroadcaster(strmID string) Broadcaster
	GetBroadcaster(strmID string) Broadcaster
	NewSubscriber(strmID string) Subscriber
	GetSubscriber(strmID string) Subscriber
}

//Broadcaster takes a streamID and a reader, and broadcasts the data to whatever underlining network.
//Note the data param doesn't have to be the raw data.  The implementation can choose to encode any struct.
//Example:
// 	s := GetStream("StrmID")
// 	b := ppspp.NewBroadcaster("StrmID", s.Metadata())
// 	for seqNo, data := range s.Segments() {
// 		b.Broadcast(seqNo, data)
// 	}
//	b.Finish()
type Broadcaster interface {
	Broadcast(seqNo uint64, data []byte) error
	Finish() error
}

//Subscriber subscribes to a stream defined by strmID.  It returns a reader that contains the stream.
//Example 1:
//	sub, metadata := ppspp.NewSubscriber("StrmID")
//	stream := NewStream("StrmID", metadata)
//	ctx, cancel := context.WithCancel(context.Background()
//	err := sub.Subscribe(ctx, func(seqNo uint64, data []byte){
//		stream.WriteSeg(seqNo, data)
//	})
//	time.Sleep(time.Second * 5)
//	cancel()
//
//Example 2:
//	sub.Unsubscribe() //This is the same with calling cancel()
type Subscriber interface {
	Subscribe(ctx context.Context, f func(seqNo uint64, data []byte)) error
	Unsubscribe() error
}

//Standard Profiles:
//1080p_60fps: 9000kbps
//1080p_30fps: 6000kbps
//720p_60fps: 6000kbps
//720p_30fps: 4000kbps
//480p_30fps: 2000kbps
//360p_30fps: 1000kbps
//240p_30fps: 700kbps
type TranscodeProfile struct {
	Name      string
	Bitrate   uint
	Framerate uint
}

type TranscodeConfig struct {
	StrmID   string
	Profiles []TranscodeProfile
}

type Transcoder interface {
	Transcode(strmID string, config TranscodeConfig, gotPlaylist func(masterPlaylist []byte)) error
}