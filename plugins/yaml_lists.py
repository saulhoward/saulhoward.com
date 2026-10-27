"""Accept YAML-style lists in Markdown metadata, so frontmatter like

    tags:
      - flood
      - games

works in both Obsidian (which wants a real YAML list) and Pelican (whose
Markdown meta parser only understands `tags: flood, games`).
"""

import re

from markdown import Markdown
from pelican import signals
from pelican.readers import MarkdownReader
from pelican.utils import pelican_open

KEY_ONLY_RE = re.compile(r"^(?P<key>[A-Za-z0-9_-]+):\s*$")
ITEM_RE = re.compile(r"^\s+-\s+(?P<item>.*?)\s*$")


def collapse_yaml_lists(text):
    lines = text.split("\n")
    start = 1 if lines and lines[0].strip() == "---" else 0
    out = lines[:start]
    i = start
    while i < len(lines):
        line = lines[i]
        if line.strip() in ("", "---", "..."):
            break  # end of metadata block
        m = KEY_ONLY_RE.match(line)
        if m:
            items = []
            while i + 1 < len(lines) and ITEM_RE.match(lines[i + 1]):
                i += 1
                items.append(ITEM_RE.match(lines[i]).group("item").strip("\"'"))
            if items:
                line = f"{m.group('key')}: {', '.join(items)}"
        out.append(line)
        i += 1
    return "\n".join(out + lines[i:])


class YamlListMarkdownReader(MarkdownReader):
    def read(self, source_path):
        self._source_path = source_path
        self._md = Markdown(**self.settings["MARKDOWN"])
        with pelican_open(source_path) as text:
            content = self._md.convert(collapse_yaml_lists(text))

        if hasattr(self._md, "Meta"):
            metadata = self._parse_metadata(self._md.Meta)
        else:
            metadata = {}
        return content, metadata


def add_reader(readers):
    for ext in YamlListMarkdownReader.file_extensions:
        readers.reader_classes[ext] = YamlListMarkdownReader


def register():
    signals.readers_init.connect(add_reader)
