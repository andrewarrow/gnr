div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2 mt-9 
      div mb-9
        {{ if eq .as "izzy" }}
          izzy is first because he's older than axl and was very much the creator, IMO.
          He drove out to the jungle a year before W. Axl Rose
        {{ else if eq .as "duff" }}
          duff is second because he was so serious about music. he carried that heavy bass and made sure gnr were gonna make it no mater what.
        {{ else if eq .as "axl" }}
          axl is next because he's axl, also but also what would gnr be without that voice.
        {{ else if eq .as "slash" }}
          I like how slash and steven ended up on the same line. Friends from middle school.
        {{ else if eq .as "steven" }}
          Steven is a great drummer. And had nothing to do with the spaghetti incident
        {{ end }}
      form id=hello flex space-x-9
        div
          img src={{.as}}.png w-32 rounded-full
        div 
          div
            textarea id=ta text-black bg-white textarea textarea-primary autofocus=true
          div mt-9 id=reply
        div 
          input id=say type=submit value=Say btn btn-primary
