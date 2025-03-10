package llm

import "github.com/openai/openai-go"

var Models = []openai.ChatModel{
	openai.ChatModelO3Mini,
	openai.ChatModelO3Mini2025_01_31,
	openai.ChatModelO1,
	openai.ChatModelO1_2024_12_17,
	openai.ChatModelO1Preview,
	openai.ChatModelO1Preview2024_09_12,
	openai.ChatModelO1Mini,
	openai.ChatModelO1Mini2024_09_12,
	openai.ChatModelGPT4_5Preview,
	openai.ChatModelGPT4_5Preview2025_02_27,
	openai.ChatModelGPT4o,
	openai.ChatModelGPT4o2024_11_20,
	openai.ChatModelGPT4o2024_08_06,
	openai.ChatModelGPT4o2024_05_13,
	openai.ChatModelGPT4oAudioPreview,
	openai.ChatModelGPT4oAudioPreview2024_10_01,
	openai.ChatModelGPT4oAudioPreview2024_12_17,
	openai.ChatModelGPT4oMiniAudioPreview,
	openai.ChatModelGPT4oMiniAudioPreview2024_12_17,
	openai.ChatModelChatgpt4oLatest,
	openai.ChatModelGPT4oMini,
	openai.ChatModelGPT4oMini2024_07_18,
	openai.ChatModelGPT4Turbo,
	openai.ChatModelGPT4Turbo2024_04_09,
	openai.ChatModelGPT4_0125Preview,
	openai.ChatModelGPT4TurboPreview,
	openai.ChatModelGPT4_1106Preview,
	openai.ChatModelGPT4VisionPreview,
	openai.ChatModelGPT4,
	openai.ChatModelGPT4_0314,
	openai.ChatModelGPT4_0613,
	openai.ChatModelGPT4_32k,
	openai.ChatModelGPT4_32k0314,
	openai.ChatModelGPT4_32k0613,
	openai.ChatModelGPT3_5Turbo,
	openai.ChatModelGPT3_5Turbo16k,
	openai.ChatModelGPT3_5Turbo0301,
	openai.ChatModelGPT3_5Turbo0613,
	openai.ChatModelGPT3_5Turbo1106,
	openai.ChatModelGPT3_5Turbo0125,
	openai.ChatModelGPT3_5Turbo16k0613,
}
