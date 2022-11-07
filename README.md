# Publish to Google Cloud Pub/Sub Topic

[![Bagde: Google Cloud](https://img.shields.io/badge/Google%20Cloud-%234285F4.svg?logo=google-cloud&logoColor=white)](#readme)
[![Bagde: Go](https://img.shields.io/badge/Go-%2300ADD8.svg?logo=go&logoColor=white)](#readme)
[![Bagde: Linux](https://img.shields.io/badge/Linux-FCC624.svg?logo=linux&logoColor=black)](#-download)
[![Bagde: macOS](https://img.shields.io/badge/macOS-000000.svg?logo=apple&logoColor=white)](#-download)
[![Bagde: Windows](https://img.shields.io/badge/Windows-008080.svg?logo=windows95&logoColor=white)](#-download)
[![Bagde: GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cyclenerd/google-cloud-pubsub-publish)](https://github.com/Cyclenerd/google-cloud-pubsub-publish/blob/master/go.mod)
[![Bagde: CI](https://github.com/Cyclenerd/google-cloud-pubsub-publish/actions/workflows/ci.yml/badge.svg)](https://github.com/Cyclenerd/google-cloud-pubsub-publish/actions/workflows/ci.yml)
[![Bagde: Release](https://github.com/Cyclenerd/google-cloud-pubsub-publish/actions/workflows/release.yml/badge.svg)](https://github.com/Cyclenerd/google-cloud-pubsub-publish/actions/workflows/release.yml)
[![Bagde: GitHub license](https://img.shields.io/github/license/cyclenerd/google-cloud-pubsub-publish)](https://github.com/Cyclenerd/google-cloud-pubsub-publish/blob/master/LICENSE)

Small (unfortunately not in file size) and fast app to publish a message to a Google Cloud Pub/Sub topic.

The app is developed in Go and faster than the original Google Cloud CLI tool `gcloud`.
There are ready-to-run compiled binaries for Linux, macOS and Windows.

## 💾 Download

<details>
<summary><b>Linux</b></summary>

Download:
* [x86_64](https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-linux-x86_64) Intel or AMD 64-Bit CPU
  ```shell
  curl -L "https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-linux-x86_64" \
       -o "pubsub-publish" && \
  chmod +x "pubsub-publish"
  ```
* [arm64](https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-linux-arm64) Arm-based 64-Bit CPU (i.e. in Raspberry Pi)
  ```shell
  curl -L "https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-linux-arm64" \
       -o "pubsub-publish" && \
  chmod +x "pubsub-publish"
  ```

To determine your OS version, run `getconf LONG_BIT` or `uname -m` at the command line.
</details>

<details>
<summary><b>macOS</b></summary>

Download:
* [x86_64](https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-macos-x86_64) Intel 64-bit
  ```shell
  curl -L "https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-macos-x86_64" \
       -o "pubsub-publish" && \
  chmod +x "pubsub-publish"
  ```
* [arm64](https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-macos-arm64) Apple silicon 64-bit
  ```shell
  curl -L "https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-macos-arm64" \
       -o "pubsub-publish" && \
  chmod +x "pubsub-publish"
  ```

To determine your OS version, run `uname -m` at the command line.
</details>

<details>
<summary><b>Windows</b></summary>

Download:
* [x86_64](https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-windows-x86_64.exe) Intel or AMD 64-Bit CPU
   ```powershell
   Invoke-WebRequest -Uri "https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-windows-x86_64.exe" -OutFile "pubsub-publish.exe"
   ```
* [arm64](https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-windows-arm64.exe) Arm-based 64-Bit CPU
   ```powershell
   Invoke-WebRequest -Uri "https://github.com/Cyclenerd/google-cloud-pubsub-publish/releases/latest/download/pubsub-publish-windows-arm64.exe" -OutFile "pubsub-publish.exe"
   ```
To determine your OS version, run `echo %PROCESSOR_ARCHITECTURE%` at the command line.
</details>

## 🔑 Authentication

Set the environment variable `GOOGLE_APPLICATION_CREDENTIALS` to the path of the JSON file that contains your service account key. This variable only applies to your current shell session, so if you open a new session, set the variable again.

<details>
<summary><b>Linux</b></summary>

Shell:

```shell
export GOOGLE_APPLICATION_CREDENTIALS="KEY_PATH"
```

Replace `KEY_PATH` with the path of the JSON file that contains your service account key.

</details>

<details>
<summary><b>macOS</b></summary>

Shell:

```shell
export GOOGLE_APPLICATION_CREDENTIALS="KEY_PATH"
```

Replace `KEY_PATH` with the path of the JSON file that contains your service account key.
</details>

<details>
<summary><b>Windows</b></summary>

PowerShell:
```powershell
$env:GOOGLE_APPLICATION_CREDENTIALS="KEY_PATH"
```

Command prompt:
```shell
set GOOGLE_APPLICATION_CREDENTIALS=KEY_PATH
```

Replace `KEY_PATH` with the path of the JSON file that contains your service account key.
</details>

## 💁 Usage

Synopsis:

```shell
pubsub-publish -project="PROJECT_ID" -topic="TOPIC" -message="MESSAGE"
```

Arguments:

**`-project`**

* Google Cloud project ID

**`-topic`**

* Topic ID or fully qualified identifier for the topic

**`-message`**

* The body of the message to publish to the given topic name


## 🚀 Benchmark

I developed this app because I wanted to publish messages faster to a Pub/Sub topic from a Raspberry Pi 4
(Broadcom BCM2711, Quad core Cortex-A72 (ARM v8) 64-bit SoC @ 1.5GHz) with Raspberry Pi OS Lite (64-bit).

Here you have the comparison of the runtime to publish one message:

| This app  |          | Google Cloud CLI |
|-----------|----------|------------------|
| 0.293 sec | +1.9 sec | 2.193 sec        |

**This app**

```shell
pi@raspberry:~ $ time ./pubsub-publish -project="PROJECT_ID" -topic="TOPIC" -message='{"test":"true"}'

real    0m0.293s
user    0m0.241s
sys     0m0.058s
```

**Google Cloud CLI**

```shell
pi@raspberry:~ $ time gcloud pubsub topics publish "TOPIC" --project="PROJECT_ID" --message='{"test":"true"}'

real    0m2.193s
user    0m1.818s
sys     0m0.256s

pi@raspberry:~ $ time gcloud version
Google Cloud SDK 401.0.0
bq 2.0.75
core 2022.09.03
gcloud-crc32c 1.0.0
gsutil 5.12

real    0m1.961s
user    0m1.722s
sys     0m0.223s
```

You can see that it is not the publication of the actual message which takes a long time,
it is the launching of the Google Cloud CLI.

## ❤️ Contributing

Have a patch that will benefit this project?
Awesome! Follow these steps to have it accepted.

1. Please read [how to contribute](CONTRIBUTING.md).
1. Fork this Git repository and make your changes.
1. Create a Pull Request.
1. Incorporate review feedback to your changes.
1. Accepted!


## 📜 License

All files in this repository are under the [Apache License, Version 2.0](LICENSE) unless noted otherwise.

Please note: No warranty
