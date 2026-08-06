#!/usr/bin/env python3
"""Regenerate the problem-index table in README.md from the repo's folder structure.

Layout on disk: <topic-folder>/<problem-folder>/submission-N.<ext>
This walks that structure and writes a markdown table (grouped by topic) between
the START/END markers in README.md.
"""
import os
from urllib.parse import quote

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
README_PATH = os.path.join(REPO_ROOT, "README.md")

START_MARKER = "<!-- PROBLEMS_INDEX:START -->"
END_MARKER = "<!-- PROBLEMS_INDEX:END -->"

IGNORED_TOP_LEVEL = {".git", ".github", "scripts"}


def encode_path(*parts):
    return "/".join(quote(p) for p in parts)


def build_index():
    topics = sorted(
        d for d in os.listdir(REPO_ROOT)
        if os.path.isdir(os.path.join(REPO_ROOT, d))
        and not d.startswith(".")
        and d not in IGNORED_TOP_LEVEL
    )

    if not topics:
        return "_No problem folders found yet._\n"

    lines = []
    total = 0
    for topic in topics:
        topic_path = os.path.join(REPO_ROOT, topic)
        problems = sorted(
            d for d in os.listdir(topic_path)
            if os.path.isdir(os.path.join(topic_path, d)) and not d.startswith(".")
        )
        if not problems:
            continue

        lines.append(f"### {topic} ({len(problems)})")
        lines.append("")
        lines.append("| # | Problem |")
        lines.append("|---|---|")
        for i, problem in enumerate(problems, start=1):
            link = encode_path(topic, problem)
            lines.append(f"| {i} | [{problem}]({link}/) |")
            total += 1
        lines.append("")

    lines.insert(0, f"_Total: {total} problems across {len(topics)} topic(s)._\n")
    return "\n".join(lines).rstrip() + "\n"


def update_readme():
    with open(README_PATH, "r", encoding="utf-8") as f:
        content = f.read()

    index_md = build_index()
    block = f"{START_MARKER}\n{index_md}{END_MARKER}"

    if START_MARKER in content and END_MARKER in content:
        start = content.index(START_MARKER)
        end = content.index(END_MARKER) + len(END_MARKER)
        new_content = content[:start] + block + content[end:]
    else:
        heading = "\n## Problem Index\n\n"
        new_content = content.rstrip("\n") + "\n" + heading + block + "\n"

    if new_content != content:
        with open(README_PATH, "w", encoding="utf-8") as f:
            f.write(new_content)
        print("README.md updated.")
    else:
        print("README.md already up to date.")


if __name__ == "__main__":
    update_readme()
