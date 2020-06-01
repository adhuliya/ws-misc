from django.shortcuts import render

# Create your views here.

def index(request):
  return render(request, "snippets/index.html", {})


def visitCount(request):
  """
  Count the visits by the client browser and show it.
  It demonstrates using the session cookies.
  Ref: https://developer.mozilla.org/en-US/docs/Learn/Server-side/Django/Sessions
  """
  # Number of visits to this view, as counted in the session variable.
  visitCount = request.session.get("visitCount", 0)
  visitCount += 1
  request.session["visitCount"] = visitCount

  # # Set session as modified to force data updates/cookie to be saved.
  # request.session.modified = True

  return render(request, "snippets/visit_count.html", {"visitCount": visitCount})


