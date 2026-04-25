#!/bin/bash

echo "🔨 Building CLI-only installer"

cd /Users/pavel/VSRPP/lab5

# Создаем директорию
mkdir -p installer/mac/build/cli

# Сборка CLI бинарника
echo "Building CLI binary..."
go build -o installer/mac/build/cli/weather ./cmd/linux/cli/main.go

chmod +x installer/mac/build/cli/weather

# Создаем README
cat > installer/mac/build/cli/README.txt << 'README'
WeatherInformer CLI

Установка:
  sudo cp weather /usr/local/bin/
  weather

Или запуск:
  ./weather

Настройка:
  ./weather -config ./config.yaml
README

# Создаем архив (переходим в папку build)
cd installer/mac/build
zip -r WeatherInformer_CLI.zip cli/

# Показываем результат
echo ""
echo "✅ CLI build complete!"
echo "📁 Location: $(pwd)/WeatherInformer_CLI.zip"
echo ""
echo "📋 Contents:"
ls -la cli/
echo ""
echo "📦 Archive size:"
ls -lh WeatherInformer_CLI.zip
