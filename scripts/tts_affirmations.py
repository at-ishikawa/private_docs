#!/usr/bin/env python3
"""Generate a single narrated audio file from the affirmations markdown.

Reads the bullet lines out of hugo/content/affirmations.md (headings, blanks and
markdown are ignored), speaks each one with the Kokoro TTS model, and inserts a
real silence after each line so you have time to say it aloud. The pause length
scales with how long the spoken line itself is.
"""
import argparse
import re
import sys

import numpy as np
import soundfile as sf
from kokoro_onnx import Kokoro


def read_affirmations(path):
    lines = []
    for raw in open(path, encoding="utf-8"):
        if raw.startswith("- "):                 # bullet lines only -> skips headings
            text = raw[2:].strip()
            text = text.replace("*", "")          # drop markdown emphasis
            text = re.sub(r"\s+", " ", text)
            if text:
                lines.append(text)
    return lines


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--src", default="hugo/content/affirmations.md")
    ap.add_argument("--model", default=".cache/kokoro/kokoro-v1.0.onnx")
    ap.add_argument("--voices", default=".cache/kokoro/voices-v1.0.bin")
    ap.add_argument("--out", default="affirmations.wav")
    ap.add_argument("--voice", default="af_heart")
    ap.add_argument("--speed", type=float, default=1.0)
    # pause after each line = spoken_line_seconds * factor + buffer
    ap.add_argument("--pause-factor", type=float, default=1.0)
    ap.add_argument("--pause-buffer", type=float, default=0.7)
    args = ap.parse_args()

    affirmations = read_affirmations(args.src)
    if not affirmations:
        sys.exit(f"No affirmations found in {args.src}")

    k = Kokoro(args.model, args.voices)
    sr = None
    chunks = []
    total_pause = 0.0

    for i, text in enumerate(affirmations, 1):
        samples, sr = k.create(text, voice=args.voice, speed=args.speed, lang="en-us")
        samples = np.asarray(samples, dtype=np.float32)
        spoken = len(samples) / sr
        pause = spoken * args.pause_factor + args.pause_buffer
        total_pause += pause
        chunks.append(samples)
        chunks.append(np.zeros(int(pause * sr), dtype=np.float32))
        print(f"[{i:>3}/{len(affirmations)}] {spoken:4.1f}s speak + {pause:4.1f}s pause  |  {text[:60]}",
              file=sys.stderr)

    audio = np.concatenate(chunks)
    sf.write(args.out, audio, sr)
    dur = len(audio) / sr
    print(f"\nWrote {args.out}: {len(affirmations)} affirmations, "
          f"{dur/60:.1f} min total ({total_pause/60:.1f} min of pauses), voice={args.voice}, speed={args.speed}",
          file=sys.stderr)


if __name__ == "__main__":
    main()
