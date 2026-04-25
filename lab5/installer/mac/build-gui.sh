#!/bin/bash

echo "🔨 Building GUI DMG installer"

cd /Users/pavel/VSRPP/lab5

APP_NAME="WeatherInformer.app"
APP_PATH="installer/mac/build/$APP_NAME"
rm -rf "$APP_PATH"
mkdir -p "$APP_PATH/Contents/MacOS"

# Собираем CLI бинарник (будет использоваться как GUI)
echo "Building application..."
go build -o "$APP_PATH/Contents/MacOS/weather" ./cmd/linux/cli/main.go

# Создаем Info.plist
cat > "$APP_PATH/Contents/Info.plist" << 'PLIST'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>weather</string>
    <key>CFBundleIdentifier</key>
    <string>com.weather.informer</string>
    <key>CFBundleName</key>
    <string>WeatherInformer</string>
    <key>CFBundleDisplayName</key>
    <string>Информер погоды</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>1.0.0</string>
    <key>CFBundleVersion</key>
    <string>1</string>
    <key>LSMinimumSystemVersion</key>
    <string>10.13</string>
</dict>
</plist>
PLIST

# Создаем DMG
DMG_PATH="installer/mac/build/WeatherInformer.dmg"
rm -f "$DMG_PATH"

DMG_DIR="installer/mac/dmg-temp"
rm -rf "$DMG_DIR"
mkdir -p "$DMG_DIR"
cp -R "$APP_PATH" "$DMG_DIR/"
ln -s /Applications "$DMG_DIR/Applications"

hdiutil create -volname "WeatherInformer" -srcfolder "$DMG_DIR" -ov -format UDZO "$DMG_PATH"

rm -rf "$DMG_DIR"

echo "✅ DMG created: $DMG_PATH"
ls -la installer/mac/build/
