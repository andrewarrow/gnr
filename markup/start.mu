div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2 mt-9
      div flex space-x-9
        div
          img src={{.as}}.png w-32 rounded-full
        div 
          textarea text-black bg-white textarea textarea-primary autofocus=true
