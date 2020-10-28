# code2img

a carbon service wrapper

## iOS Shortcut

Basic usage notes:

- Get the shortcut from here: https://www.icloud.com/shortcuts/6ac8f9afc47e4b109bca81648c59e2f4
- The shortcut also writes a URL back to your clipboard, you can paste the link to your browser for a better copy experience.
- **Remember: Keep your life simple. Keep in mind that you do not upload more than 50 lines of code. Otherwise, no one wants to read it :)**
<!-- ffmpeg -i record.mp4 -vf scale=288:640 demo.gif -->

Demo:

![](./demo.gif)

## API

```
POST golang.design/api/v1/code2img
{
    "code": "code string"
}
```

Response pure text (better for iOS shortcut):

```
https://golang.design/api/v1/code2img/data/images/06ad29c5-2989-4a8e-8cd2-1ce63e36167b.png
```

You can also access the code text:

```
https://golang.design/api/v1/code2img/data/code/06ad29c5-2989-4a8e-8cd2-1ce63e36167b.go
```

## License

&copy; 2020 The golang.design Initiative Authors.