from django.shortcuts import render
import logging

logger = logging.getLogger("mysite.custom")

# Create your views here.

def index(request):
  context = {}
  return render(request, 'main/index.html', context)


def errorHandler(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {}
  return render(request, 'main/error_page.html', context)


def handler404(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": 404}
  return render(request, 'main/error_page.html', context)


def handler403(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": 403}
  return render(request, 'main/error_page.html', context)


def handler500(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": 500}
  return render(request, 'main/error_page.html', context)


