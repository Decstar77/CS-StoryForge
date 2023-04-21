# CS-StoryForge

## Introduction

Welcome to the CS-StoryForge repository! This small program is designed to generate a prompt based on a Counter-Strike: Global Offensive (CSGO) demo file. Once the prompt is generated, you can send it to GPT, which will create an engaging story and summary based on the gameplay.

See [example outputs](./example_outputs) from professional games

Please keep in mind it is currently in its first version (v1.0). As such, it may contain bugs and other issues.

## Installation and Usage

Download the pre-built .exe from the releases section. Upon running the program, a file explorer window will open for you to select your demo file. After selection, the program will generate a .txt file with the prompt. You can then copy this prompt into a language model like GPT.

## Tips

Changing the first sentence of the prompt can have drastic changes on the output. Here are some techniques for better output

1. Add a specific player, e.g., "Write a story about a CSGO match highlighting s1mple."
2. Specify the game time, e.g., "Focus on the end game."
3. Adjust adjectives to change the tone.

## Building
1. Ensure that you have Golang installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the cloned repository and run the following command to build the binary: `go build`
4. Upon successful compilation, you will see a binary named `CS-StoryForge` in your current directory.

## Future
Some thoughts on how to improve the prompt.
1. Information on locations and player movements to provide greater context and facilitate a deeper understanding of the game.
2. Updates on team/player wealth after each round, hopefully offering insights into their financial status and potential strategies.
3. Eliminate mundane rounds, such as uneventful eco rounds, to maintain the LLM's interest on pivotal moments while also reducing token count.
4. Include screenshots, such as of the mini map or perspectives from key players.

Another approach could be to generate a much more comprehensive and detailed prompt for each round. This could lead to short stories on certain rounds rather than a board overview. Explore the possibility of narrating the events in the style of a specific caster and utilizing AI-generated voices.

## License
This project is released under the MIT License. See [LICENSE](LICENCE) for more information.

## Acknowledgements
A huge thank you to the community for your support, feedback, and contributions!

