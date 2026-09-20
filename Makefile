SHELL := /bin/bash

# Pauses are built by repeating [long pause] tags to cover the time it takes to
# speak each line aloud. SPEAK_CPS = how fast you speak (characters/second);
# LONG_PAUSE_SECS = seconds one [long pause] tag lasts in your voice AI.
# Override e.g.: `make affirmations SPEAK_CPS=10` (more time per line).
SPEAK_CPS ?= 13
LONG_PAUSE_SECS ?= 1.5

# `make audio` settings. VOICE: any Kokoro voice (af_heart, af_bella, am_michael,
# bm_george, ...). PAUSE_FACTOR/PAUSE_BUFFER: gap after each line = spoken length
# x factor + buffer seconds. Override e.g.: `make audio VOICE=am_michael PAUSE_FACTOR=1.3`.
VOICE ?= af_heart
SPEED ?= 1.0
PAUSE_FACTOR ?= 1.0
PAUSE_BUFFER ?= 2.7
AUDIO_OUT ?= affirmations.wav

.PHONY: start affirmations audio

start:
	hugo server --buildDrafts --source hugo --bind 0.0.0.0

# Generate a single narrated audio file (Kokoro TTS, fully local/offline) with a
# real silence after each affirmation sized to how long the line takes to speak.
audio:
	.venv-tts/bin/python scripts/tts_affirmations.py \
		--src hugo/content/affirmations.md \
		--model .cache/kokoro/kokoro-v1.0.onnx \
		--voices .cache/kokoro/voices-v1.0.bin \
		--out $(AUDIO_OUT) --voice $(VOICE) --speed $(SPEED) \
		--pause-factor $(PAUSE_FACTOR) --pause-buffer $(PAUSE_BUFFER)
