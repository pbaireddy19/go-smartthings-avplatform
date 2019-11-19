package models

import "time"

type GetSourceResponse struct {
	Source struct {
		ID       string `json:"id"`
		UserID   string `json:"user_id"`
		ClientID string `json:"client_id"`
		Name     string `json:"name"`
		Profiles struct {
			High struct {
				MaxClipDuration int      `json:"max_clip_duration"`
				Groups          []string `json:"groups"`
				Media           struct {
					BitRate int `json:"bit_rate"`
					Tracks  []struct {
						Type       string `json:"type"`
						Codec      string `json:"codec"`
						Width      int    `json:"width,omitempty"`
						Height     int    `json:"height,omitempty"`
						FrameRate  int    `json:"frame_rate,omitempty"`
						BitRate    int    `json:"bit_rate"`
						SampleRate int    `json:"sample_rate,omitempty"`
					} `json:"tracks"`
				} `json:"media"`
			} `json:"high"`
			Medium struct {
				MaxClipDuration int      `json:"max_clip_duration"`
				Groups          []string `json:"groups"`
				Media           struct {
					BitRate int `json:"bit_rate"`
					Tracks  []struct {
						Type       string `json:"type"`
						Codec      string `json:"codec"`
						Width      int    `json:"width,omitempty"`
						Height     int    `json:"height,omitempty"`
						FrameRate  int    `json:"frame_rate,omitempty"`
						BitRate    int    `json:"bit_rate"`
						SampleRate int    `json:"sample_rate,omitempty"`
					} `json:"tracks"`
				} `json:"media"`
			} `json:"medium"`
			Low struct {
				MaxClipDuration int      `json:"max_clip_duration"`
				Groups          []string `json:"groups"`
				Media           struct {
					BitRate int `json:"bit_rate"`
					Tracks  []struct {
						Type       string `json:"type"`
						Codec      string `json:"codec"`
						Width      int    `json:"width,omitempty"`
						Height     int    `json:"height,omitempty"`
						FrameRate  int    `json:"frame_rate,omitempty"`
						BitRate    int    `json:"bit_rate"`
						SampleRate int    `json:"sample_rate,omitempty"`
					} `json:"tracks"`
				} `json:"media"`
			} `json:"low"`
		} `json:"profiles"`
		DefaultProfile string    `json:"default_profile"`
		Location       string    `json:"location"`
		Created        time.Time `json:"created"`
		Modified       time.Time `json:"modified"`
		Online         bool      `json:"online"`
		LastActive     time.Time `json:"last_active"`
		MotionEnabled  bool      `json:"motion_enabled"`
		Disabled       bool      `json:"disabled"`
		Properties     struct {
			DeviceTypeName    string `json:"device_type_name"`
			Continent         string `json:"continent"`
			Country           string `json:"country"`
			ClipRetention     int    `json:"clip_retention"`
			ClipDurationLimit int    `json:"clip_duration_limit"`
			RSSI              int    `json:"RSSI"`
			SignalStrength    int    `json:"signal_strength"`
			RoomID            string `json:"room_id"`
		} `json:"properties"`
		ClipDuration    int      `json:"clip_duration"`
		SerialNumber    string   `json:"serial_number"`
		FirmwareVersion string   `json:"firmware_version"`
		Manufacturer    string   `json:"manufacturer"`
		Vendor          string   `json:"vendor"`
		Model           string   `json:"model"`
		ActiveEntities  string   `json:"active_entities"`
		Features        []string `json:"features"`
		CustomSettings  struct {
			IrLight               string `json:"ir_light"`
			SpeakerVolume         int    `json:"speaker_volume"`
			StatusLight           bool   `json:"status_light"`
			FlipImage             bool   `json:"flip_image"`
			MirrorImage           bool   `json:"mirror_image"`
			Watermark             bool   `json:"watermark"`
			Hdr                   bool   `json:"hdr"`
			LensDewarping         bool   `json:"lens_dewarping"`
			MuteClipAudio         bool   `json:"mute_clip_audio"`
			SoundEventSensitivity int    `json:"sound_event_sensitivity"`
			DemoIP                string `json:"demo_ip"`
		} `json:"custom_settings"`
		ClipRetentionDays int `json:"clip_retention_days"`
	} `json:"source"`
}
