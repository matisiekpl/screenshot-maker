# Screenshot Maker

Generate fancy screenshots for AppStore and Google Play from original screenshot and text.

<div style="text-align: center">
  <img src="https://github.com/matisiekpl/screenshot-maker/assets/21008961/c1dcaf33-3d42-41d4-a687-cea96733c140" width="30%">
</div>

# Install
```bash
go install github.com/matisiekpl/screenshot-maker@latest
```

# Usage
```bash
screenshot-maker --input originalScreenshot.png --text "Some text"
```

You can also set device. 
Available devices:
- `iphone-11`
- `iphone-11-alt`
- `ipad-pro`

Setting `--device` generates diffrent output files. Dimensions are adjusted for AppStore Guidelines. After generation, you can directly upload output files to AppStore.

For example for generating `ipad-pro` (size guidelines: `2048x2732`):
```bash
screenshot-maker --input originalScreenshot.png --text "Some text" --device ipad-pro
```
