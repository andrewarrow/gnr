say -v Eddy "Now that I have AI doing sign language with insects, what should I do with this technology? That is your question? Silly human. You know the answer. Fight the real enemy? Do you even know your enemy?" -o sign.aiff

#ffmpeg -f lavfi -i color=c=black:s=1280x720:d=30 -r 30 -pix_fmt yuv420p black.mp4\n

ffmpeg -i black.mp4 -i sign.aiff -c:v copy -c:a aac -shortest black_with_audio.mp4\n
ffmpeg -i black_with_audio.mp4 -vf subtitles=subtitles.srt -c:a copy output.mp4\n
