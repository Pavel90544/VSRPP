.PHONY: help installer-cli installer-gui installer-clean

help:
	@echo "Available commands:"
	@echo "  make installer-cli  - Build CLI only installer (ZIP)"
	@echo "  make installer-gui  - Build GUI DMG installer"
	@echo "  make installer-clean - Clean installers"

installer-cli:
	@./installer/mac/build-cli.sh

installer-gui:
	@./installer/mac/build-gui.sh

installer-clean:
	@rm -rf installer/mac/build/*
	@echo "Clean complete!"
