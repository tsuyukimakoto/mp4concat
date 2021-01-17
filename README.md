# mp4concat

Files shot with GoPro are split into 4.01GB pieces.

You can combine multiple files that have been split into one file without re-encoding (no video degradation).

The ffmpeg command is required.

## How to use

Install ffmpeg beforehand.

Double-click on mp4concat to start the terminal.

Display the contents of your GoPro's SD card in order of oldest to newest, and select all the video files you want to merge into one video. Only MP4 files will be used.

Once you have pasted the list of files, activate the terminal (click on it) and hit the enter key. If you expand the window a little, you can see the status without scrolling.

## What's happening

Create a file in the ~/Desktop/mp4concat folder with a list of the files passed to you (in the format used by ffmpeg's concat).

## Platforms that are likely to work

I tried to run it on macOS 11.1 (Big Sur).

go: 1.15

---

GoProで撮影したファイルは4.01GBごとに分割されてしまいます。

分割されてしまった複数ファイルを、再エンコードを行わずに1つのファイルに結合します（動画の劣化がありません）。

ffmpegコマンドが必要です。

## 使い方

あらかじめffmpegをインストールしておきます。

mp4concatをダブルクリックするとterminalが起動します。

GoProの撮影済みのSDカードの中身を日付の古い順に表示して、1つの動画に結合したい動画ファイルを全て選択します。このとき、THMファイルも一緒に選択してしまって構いません。MP4ファイルだけを利用します。

ファイルの一覧を貼り付けられたら、terminalをアクティブにして（クリックして）、エンターキーを推します。ウィンドウを少し広げるとスクロールせずに状況が表示されます。

## 何が起きているのか

~/Desktop/mp4concat フォルダに、渡されたファイルの一覧が記載されたファイルで作成します（ffmpegのconcatで使う形式で）。

## 動作すると思われるプラットフォーム

macOSの11.1(Big Sur)で動作させてみました。

go: 1.15