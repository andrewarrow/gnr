say -v Thomas "hard code it; each shouldUseAroma function says 1/10 random calls is not webscale" -o webscale.aiff


  ffmpeg -f lavfi -i color=c=black:s=1280x720:d=30 -r 30 -pix_fmt yuv420p black.mp4\n
  ffmpeg -i black.mp4 -i webscale.aiff -c:v copy -c:a aac -shortest black_with_audio.mp4\n
  histor
  history
  vi sub.txt
  vi sub.txt
  mv sub.txt subtitles.srt
  ffmpeg -i black_with_audio.mp4 -vf subtitles=subtitles.srt -c:a copy output.mp4\n
  open output.mp4
