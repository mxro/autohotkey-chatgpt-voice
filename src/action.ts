import * as express from "express";
import * as fs from "fs";
import * as ffmpeg from "fluent-ffmpeg";
import { Readable } from "stream";
// import ffmpegPath from "@ffmpeg-installer/ffmpeg";

const app = express();
const port = 10101;

// // Configure OpenAI API
// const configuration = new Configuration({
//   apiKey: "YOUR_OPENAI_API_KEY",
// });
// const openai = new OpenAIApi(configuration);

// Set ffmpeg path
// ffmpeg.setFfmpegPath(ffmpegPath);

// Record audio
function recordAudio(filename: string): Promise<void> {
  return new Promise((resolve, reject) => {

    var ai = new portAudio.AudioIO({
    inOptions: {
      channelCount: 2,
      sampleFormat: portAudio.SampleFormat16Bit,
      sampleRate: 44100,
      deviceId: -1, // Use -1 or omit the deviceId to select the default device
      closeOnError: true // Close the stream if an audio error is detected, if set false then just log the error
    }
  });

  // Create a write stream to write out to a raw audio file
  var ws = fs.createWriteStream('rawAudio.raw');

  //Start streaming
  ai.pipe(ws);
  ai.start();

    // Start recording on POST request
    app.post("/", async (req, res) => {
      if (req.body === "endRecording") {
        console.log("POST request received to end recording");
        const audioFilename = "recorded_audio.wav";
        // const transcription = await transcribeAudio(audioFilename);
        // console.log("Transcription:", transcription);
        // await convertToMp3(audioFilename);
        res.sendStatus(200);
        process.exit(0);
      } else {
        res.sendStatus(400);
      }
    });

    // Start the server
    app.listen(port, () => {
      console.log(`Server is listening on port ${port}`);
    });

    process.on("SIGINT", () => {
      micInstance.stop();
      console.log("Finished recording");
      resolve();
    });

    micInputStream.on("error", (err) => {
      reject(err);
    });
  });
}

// Transcribe audio
// async function transcribeAudio(filename: string): Promise<string> {
//   const transcript = await openai.createTranscription(
//     fs.createReadStream(filename),
//     "whisper-1"
//   );
//   return transcript.data.text;
// }

// Convert audio to mp3
async function convertToMp3(filename: string): Promise<void> {
  return new Promise((resolve, reject) => {
    ffmpeg(filename)
      .toFormat("mp3")
      .on("end", () => {
        console.log("Conversion to mp3 completed");
        resolve();
      })
      .on("error", (err) => {
        reject(err);
      })
      .save(`${filename}.mp3`);
  });
}


// Main function
async function main() {
  const audioFilename = "recorded_audio.wav";
  await recordAudio(audioFilename);
  // const transcription = await transcribeAudio(audioFilename);
  // console.log("Transcription:", transcription);
}

main();
