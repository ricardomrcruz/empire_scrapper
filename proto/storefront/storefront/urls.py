from django.conf import settings
from django.conf.urls.static import static
from django.contrib import admin
from django.urls import path, include

from playground.views import index, contact


# hello
urlpatterns = [
    path("", index, name="index"),
    path("contact/", contact, name="contact"),
    path("admin/", admin.site.urls),
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)
