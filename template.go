package main

var indexTemplate = `<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <script src="https://cdn.tailwindcss.com"></script>
  <title>go-ram</title>
</head>

<body>
<div class="w-screen p-24 text-center h-screen bg-gradient-to-br from-red-400 to-yellow-600 grid place-items-center ">
  <p class="text-6xl font-semibold">{{.Insult}}</p>
</div>
</body>

</html>`
