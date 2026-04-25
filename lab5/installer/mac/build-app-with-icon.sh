#!/bin/bash

echo "🔨 Building WeatherInformer.app with custom icon"

cd /Users/pavel/VSRPP/lab5

APP_NAME="WeatherInformer.app"
APP_PATH="installer/mac/build/$APP_NAME"
rm -rf "$APP_PATH"
mkdir -p "$APP_PATH/Contents/MacOS"
mkdir -p "$APP_PATH/Contents/Resources"

# Копируем иконку
if [ -f "assets/icons/WeatherInformer.icns" ]; then
    cp assets/icons/WeatherInformer.icns "$APP_PATH/Contents/Resources/"
    echo "✅ Icon copied"
else
    echo "⚠️ Icon not found, building without icon"
fi

# Собираем CLI бинарник
echo "Building application binary..."
go build -o "$APP_PATH/Contents/MacOS/weather" ./cmd/linux/cli/main.go

# Создаем запускающий скрипт
cat > "$APP_PATH/Contents/MacOS/WeatherInformer" << 'SCRIPT'
#!/bin/bash
DIR="$(cd "$(dirname "$0")" && pwd)"
CONFIG_DIR="$HOME/.weather"
if [ ! -f "$CONFIG_DIR/config.yaml" ]; then
    mkdir -p "$CONFIG_DIR"
    cat > "$CONFIG_DIR/config.yaml" << 'CONFIG'
service:
  provider:
    type: open-meteo
  location:
    lat: 53.9045
    long: 27.5615
  cache:
    type: memory
    ttl: 300
CONFIG
fi
exec "$DIR/weather" -config "$CONFIG_DIR/config.yaml"
SCRIPT

chmod +x "$APP_PATH/Contents/MacOS/WeatherInformer"

# Создаем Info.plist
cat > "$APP_PATH/Contents/Info.plist" << 'PLIST'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>WeatherInformer</string>
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
    <key>NSHighResolutionCapable</key>
    <true/>
</dict>
</plist>
PLIST

# Создаем DMG
echo "Creating DMG installer..."
DMG_PATH="installer/mac/build/WeatherInformer.dmg"
rm -f "$DMG_PATH"

DMG_DIR="installer/mac/dmg-temp"
rm -rf "$DMG_DIR"
mkdir -p "$DMG_DIR"
cp -R "$APP_PATH" "$DMG_DIR/"
ln -s /Applications "$DMG_DIR/Applications"

hdiutil create -volname "WeatherInformer" -srcfolder "$DMG_DIR" -ov -format UDZO "$DMG_PATH"

rm -rf "$DMG_DIR"

echo ""
echo "✅ App built: $APP_PATH"
echo "✅ DMG created: $DMG_PATH"
echo ""
echo "📋 To install: open $DMG_PATH and drag to Applications"
