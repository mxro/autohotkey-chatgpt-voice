"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
var express = require("express");
var fs = require("fs");
var ffmpeg = require("fluent-ffmpeg");
var mic_1 = require("mic");
var stream_1 = require("stream");
// import ffmpegPath from "@ffmpeg-installer/ffmpeg";
var app = express();
var port = 10101;
// // Configure OpenAI API
// const configuration = new Configuration({
//   apiKey: "YOUR_OPENAI_API_KEY",
// });
// const openai = new OpenAIApi(configuration);
// Set ffmpeg path
// ffmpeg.setFfmpegPath(ffmpegPath);
// Record audio
function recordAudio(filename) {
    var _this = this;
    return new Promise(function (resolve, reject) {
        var micInstance = (0, mic_1.default)({
            rate: "16000",
            channels: "1",
            fileType: "wav",
        });
        var micInputStream = micInstance.getAudioStream();
        var output = fs.createWriteStream(filename);
        var writable = new stream_1.Readable().wrap(micInputStream);
        console.log("Recording... Press Ctrl+C to stop.");
        writable.pipe(output);
        micInstance.start();
        // Start recording on POST request
        app.post("/", function (req, res) { return __awaiter(_this, void 0, void 0, function () {
            var audioFilename;
            return __generator(this, function (_a) {
                if (req.body === "endRecording") {
                    console.log("POST request received to end recording");
                    audioFilename = "recorded_audio.wav";
                    // const transcription = await transcribeAudio(audioFilename);
                    // console.log("Transcription:", transcription);
                    // await convertToMp3(audioFilename);
                    res.sendStatus(200);
                    process.exit(0);
                }
                else {
                    res.sendStatus(400);
                }
                return [2 /*return*/];
            });
        }); });
        // Start the server
        app.listen(port, function () {
            console.log("Server is listening on port ".concat(port));
        });
        process.on("SIGINT", function () {
            micInstance.stop();
            console.log("Finished recording");
            resolve();
        });
        micInputStream.on("error", function (err) {
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
function convertToMp3(filename) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            return [2 /*return*/, new Promise(function (resolve, reject) {
                    ffmpeg(filename)
                        .toFormat("mp3")
                        .on("end", function () {
                        console.log("Conversion to mp3 completed");
                        resolve();
                    })
                        .on("error", function (err) {
                        reject(err);
                    })
                        .save("".concat(filename, ".mp3"));
                })];
        });
    });
}
// Main function
function main() {
    return __awaiter(this, void 0, void 0, function () {
        var audioFilename;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    audioFilename = "recorded_audio.wav";
                    return [4 /*yield*/, recordAudio(audioFilename)];
                case 1:
                    _a.sent();
                    return [2 /*return*/];
            }
        });
    });
}
main();
