AUTHOR = "Saul Howard"
SITENAME = "Saul Howard"

THEME = "theme"

PATH = "content"

TIMEZONE = "Europe/London"

DEFAULT_LANG = "en"

ARTICLE_URL = "{slug}"
PAGE_URL = "{slug}"
PAGE_SAVE_AS = "{slug}.html"
ARCHIVES_SAVE_AS = "posts.html"

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None

DISPLAY_CATEGORIES_ON_MENU = False
DISPLAY_PAGES_ON_MENU = False

MENUITEMS = (
    ("Home", "/"),
    ("Posts", "/posts"),
    ("Talks & Press", "/talks"),
)

DEFAULT_PAGINATION = 10

STATIC_PATHS = ["images", "extra/CNAME", "extra/brick-favicon.ico", "extra/brick-favicon.svg"]
EXTRA_PATH_METADATA = {
    "extra/CNAME": {"path": "CNAME"},
    "extra/brick-favicon.ico": {"path": "brick-favicon.ico"},
    "extra/brick-favicon.svg": {"path": "brick-favicon.svg"},
}
IGNORE_FILES = [".Ulysses-Group.plist"]

# Uncomment following line if you want document-relative URLs when developing
# RELATIVE_URLS = True

PLUGIN_PATHS = ["plugins"]
PLUGINS = ["yaml_lists"]
