<!DOCTYPE html>
<html>
<head lang="{{lang}}">
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="robots" content="noindex, nofollow">
    <title>{{page_title | default:"Admin" }}</title>
</head>
<body>

{% block body %}
{% block content %}{% endblock %}
{% endblock %}

</body>
</html>