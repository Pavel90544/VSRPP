#!/bin/bash
# Создаем квадрат 1024x1024
convert -size 1024x1024 xc:"#2196F3" icon_1024.png 2>/dev/null || \
sips -z 1024 1024 --setProperty format png -s formatOptions 70 /System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/GenericFolderIcon.icns --out icon_1024.png 2>/dev/null
