Bot assumptions: This is the first implementation of the bot, so we are making some assumptions about how it will work and keeping it small:
- Bot only trades DogBones.
- Bot buys when price is low.
- Bot sells when price is high.
- Bot does nothing when price is normal.
- First version is deterministic: same price always gives same bot action.
- No randomness is used.