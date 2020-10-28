# code2img

a carbon service wrapper

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

## iOS Shortcut

<!-- ffmpeg -i record.mp4 -vf scale=288:640 demo.gif -->

![](./demo.gif)

Get the shortcut from here: https://www.icloud.com/shortcuts/dac1a52db1d64cd79b5baaacf262fb5b

**Remember: Do not upload code longer than 50 lines. Keep your life simple.**

## License

&copy; 2020 The golang.design Initiative Authors.