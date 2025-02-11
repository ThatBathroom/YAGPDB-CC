{{ $ban_emoji := "🔨" }}
{{ $timeout_emoji := "⏰" }}
{{ $kick_emoji := "🚪" }}
{{ $channelID := Replace with channelID }}

{{ if eq .Reaction.Emoji.Name $ban_emoji }}
  {{ $authorID := .Message.Author.ID }}
  {{ $reason := "yep, buddy your banned" }}
  {{ print "Banning Author ID: " $authorID }}
  {{ execAdmin "ban" $authorID $reason }}
  {{ sendMessage $channelID (print "Author ID: " $authorID " has been banned. Reason: " $reason) }}
{{ end }}

{{ if eq .Reaction.Emoji.Name $timeout_emoji }}
  {{ $authorID := .Message.Author.ID }}
  {{ $duration := "10m" }}
  {{ $reason := "went out of way to break rules" }}
  {{ print "Timing out Author ID: " $authorID " for " $duration }}
  {{ execAdmin "timeout" $authorID $duration $reason }}
  {{ sendMessage $channelID (print "Author ID: " $authorID " has been timed out for " $duration ". Reason: " $reason) }}
{{ end }}

{{ if eq .Reaction.Emoji.Name $kick_emoji }}
  {{ $authorID := .Message.Author.ID }}
  {{ $reason := "sometimes I just like kicking people" }}
  {{ print "Kicking Author ID: " $authorID }}
  {{ execAdmin "kick" $authorID $reason }}
  {{ sendMessage $channelID (print "Author ID: " $authorID " has been kicked. Reason: " $reason) }}
{{ end }}
