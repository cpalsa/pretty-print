#!/bin/bash

# load clipboard 
clipboard=$(xclip -o -selection clipboard)

# Detect clipboard format
if [[ -z "$clipboard" || "$clipboard" =~ ^\s*$ ]]; then
  # Clipboard is empty or contains only whitespace
  notify-send "Clipboard does not contain valid XML or JSON."
elif [[ "$clipboard" =~ ^\".*\"$ ]] || [[ "$clipboard" =~ ^\'.*\'$ ]]; then
  # Clipboard contains a quoted string
  notify-send "Clipboard does not contain valid XML or JSON."
elif xmlstarlet fo -t <<< "$clipboard" &> /dev/null; then
  # Format clipboard contents as XML and update clipboard
  xmlstarlet fo -t <<< "$clipboard" | sed '1d' | xclip -i -selection clipboard
  notify-send "XML formatting completed successfully!"
elif jq '.' <<< "$clipboard" &> /dev/null; then
  # Format clipboard contents as JSON and update clipboard
  jq '.' <<< "$clipboard" | xclip -i -selection clipboard
  notify-send "JSON formatting completed successfully!"
else
  # Clipboard does not contain valid XML or JSON
  notify-send "Clipboard does not contain valid XML or JSON."
fi
