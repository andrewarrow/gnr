say -v Eddy "Now that I have AI doing sign language with insects, what should I do with this technology? That is your question? Silly human. You know the answer. Fight the real enemy? Do you even know your enemy?" -o sign.aiff

#ffmpeg -f lavfi -i color=c=black:s=1280x720:d=30 -r 30 -pix_fmt yuv420p black.mp4\n

ffmpeg -i black.mp4 -i sign.aiff -c:v copy -c:a aac -shortest black_with_audio.mp4\n
ffmpeg -i black_with_audio.mp4 -vf subtitles=subtitles.srt -c:a copy output.mp4\n


ffmpeg -i input_video.mp4 -an -c:v copy sign2.mp4

ffmpeg -i french.aiff -i french.m4a -filter_complex "[0:a]volume=1.0[a0];[1:a]volume=0.5[a1];[a0][a1]amix=inputs=2:duration=longest" mixed.mp3

sox -m french.aiff french.m4a mixed.mp3 vol 1,0.5
