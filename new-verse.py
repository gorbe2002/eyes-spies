import openai
openai.api_key = ""

# Ice Spice Verse Generator:
song1 = ("""
Song Name: Munch (Feelin' U)

[Intro]
Grrah
(Stop playin' with 'em, RIOT)
Grrah

[Chorus]
You thought I was feelin' you? (Nah)
That nigga a munch
Nigga a eater, he ate it for lunch
Bitch, I'm a baddie, I get what I want, like
You thought I was feelin' you? (Nah, thought I was feelin' you?)
That nigga a munch
Nigga a eater, he ate it for lunch
Bitch, I'm a baddie, I get what I want, like (Get what I want, like)

[Verse 1]
Bitches ain't bad, let's keep it a bean
Know they be mad that I be on the scene
Ass too fat, can't fit in no jeans
You was my stitch but it's not what it seam
I got that wetty, I'm keepin' it clean
Fuckin' with niggas that's totin' a beam (Grrah)
Sayin' you love me but what do you mean? (Grrah)
Pretty as fuck and he like that I'm mean
Baddest bitch out you shittin' me? (Shittin' me)
If you ain't a baddie can't sit with me (Sit with me)
I swear that these bitches my mini-me's (Mini-me, grrah, grrah)
He wanna sex, niggas be dreamin'
I'm from the X, niggas be schemin'
I'm on they necks, they is not breathin'
Thumbin' a check, blow it in Neiman's

[Chorus]
You thought I was feelin you? (Nah)
That nigga a munch
Nigga a eater, he ate it for lunch
Bitch, I'm a baddie I get what I want, like
You thought I was feelin you? (Nah)
That nigga a munch
Nigga a eater, he ate it for lunch
Bitch, I'm a baddie I get what I want, like (Get what I want, like)

[Verse 2]
Bitch, I'm a baddie I get what I please
You know my body, I do it with ease
He want my body, he tellin' me, "Please" (Grrah)
I'm walkin' past him, he sniffin' my breeze
He jackin' me, but he not my boo (Grrah)
He like the jewelry I wear on my boobs
How can I link you when I got a shoot? (Huh?)
Don't want your love, I just want the blue (Huh?)
Grabbin' my ass while I'm doin' my dance
She keep on starin' 'cause shorty a fan (Damn)
Gotta stick to the plan
He mad as fuck, I won't give him a chance
But still he gon' do what I say (Do what I say)
I swear I be stuck in my ways (Stuck in my ways)
But still he gon' do what I say (Do what I say)
I swear I be stuck in my ways (Grrah)

[Chorus]
You thought I was feelin you? (Nah)
That nigga a munch
Nigga a eater, he ate it for lunch
Bitch, I'm a baddie I get what I want, like
You thought I was feelin you? (Nah)
That nigga a munch
Nigga a eater, he ate it for lunch
Bitch, I'm a baddie I get what I want, like
""")

song2 = ("""
Song Name: In Ha Mood

[Intro]
Where were you last week, when you stopped coming by?
Stop playin' with 'em, RIOT

[Chorus]
Like, damn, she in her mood (Grrah)
Like, damn, she in her mood (Mood)
Like, damn, she in her mood (In her mood, she in her mood)
Like, damn, she in her mood (She in her mood)
She lit, get money too (Like)
Like, damn, she in her mood (She in her mood), damn

[Verse 1]
In the mirror, I'm doin' my dance (Like)
And he packin', I know by his pants (Grrah)
He a rapper, but don't got a chance
Stuck in my ways so I'm lovin' my bands (Damn)
Like a million views in a day (Like)
It's so many ways to get paid (Grrah)
I tried dippin', he begged me to stay
Bae, I'm not stayin', I just wanna play (Just wanna play)
In the party, he just wanna rump (Rump)
Big boobs and the butt stay plump (Stay plump)
She a baddie, she know she a ten (Baddie, ten)
She a baddie with her baddie friend (Damn, friend)
They like, "Ice, how you always stay hot?" (Hot)
Oh, they mad 'cause I keep makin' bops (Bops)
Oh, she mad 'cause I'm takin' her spot
If I was bitches, I'd hate me a lot (Grrah)

[Chorus]
Like, damn, she in her mood (Grrah)
Like, damn, she in her mood (Mood)
Like, damn, she in her mood (In her mood, she in her mood)
Like, damn, she in her mood (She in her mood)
She lit, get money too (Like)
Like, damn, she in her mood (She in her mood), damn

[Verse 2]
No friends, I don't fuck with the fakes (Grrah)
Sayin' they love me, but wantin' my place (Like)
Step in the party, I'm lookin' the baddest
So the paparazzi in my face (Grrah)
Pretty bitch, but I came from the gutter
Said I'd be lit by the end of the summer (Like)
And I'm proud that I'm still gettin' bigger (Damn)
Goin' viral is gettin' 'em sicker
Like, what? Let's keep it a buck (Huh)
Bitches too borin', got 'em stuck in a rut (Damn)
Lamborghini roarin' when I hop out the truck (Huh)
Pretty bitch like Lauryn with a big ass butt, yup
Pretty face and the waist all gone (Huh)
And I'm makin' 'em wait, hold on (Hold on)
And I'm makin' 'em wait, hold on (Hold on)
Wait, hold on (Grrah, hold on)

[Chorus]
Like, damn, she in her mood (Grrah)
Like, damn, she in her mood (Mood)
Like, damn, she in her mood (In her mood, she in her mood)
Like, damn, she in her mood (She in her mood)
She lit, get money too (Like)
Like, damn, she in her mood (She in her mood, damn)
Like, damn, she in her mood (Grrah)
Like, damn, she in her mood (Mood)
Like, damn, she in her mood (In her mood, she in her mood)
Like, damn, she in her mood (She in her mood)
She lit, get money too (Like)
Like, damn, she in her mood (She in her mood, damn)
""")

song3 = ("""
Song Name: Bikini Bottom

[Intro]
(Stop playin' with 'em, RIOT)

[Chorus]
How can I lose if I'm already chose? Like
If she feelin' hot, then I make that bitch froze
And I get a bitch tight every time that I post, damn
The party not lit, then I'd rather not go (Why would I go?)
If she feelin' hot, then I make that bitch froze
And I get a bitch tight every time that I post, damn

[Verse 1]
Bitches down bad, they be on my ass
I can hear you hatin' from the back (The back)
Balenciaga baddie, got a bag (A bag)
Nigga munchin', ate it from the back (The back)
Nigga fiendin', gotta play it cool (Huh?)
Got the jatty, I'ma make it move (Goddamn)
Breakin' records and I'm breakin' news
Bitches be pressed, like, "Who you?" (Who you?)
I get whatever I like
Bitches won't bark, but they wanna bite (Grah)
I got two milli' for usin' a mic, bitch
Think about that when you type (Haha)
She wanna party with Spice (Party with Spice)
And the body gon' eat, bon appétit
Ass on fat with the waist on sleek (Yeah)
Ginger hair, pretty, calm, beat

[Chorus]
How can I lose if I'm already chose? Like
If she feelin' hot, then I make that bitch froze
And I get a bitch tight every time that I post, damn
The party not lit, then I'd rather not go (Why would I go?)
If she feelin' hot, then I make that bitch froze
And I get a bitch tight every time that I post, damn

[Verse 2]
Look in the mirror, I'm feelin' me (Feelin' me)
Stackin' money where the ceilin' be (Ceilin' be)
Twenty to stand on a couch (Couch)
When out of town, fuck if they feelin' me (Feelin' me)
He like the way that I dress (Yes)
Throw on Balenci' to flex (Grrah)
Like MiMi, I got 'em obsessed (Obsessed)
Them bitches see me as a threat (Yes)
The baddest in the room (Baddest in the room)
So tell 'em to make room
Diamonds glisten on my boob (Boob)
They gon' listen to my tune (Tune)
Flow the nicest, but I'm rude (Rude)
I like niggas, bitches too (Bitches too)
Ayo, baddie, what it do? (What it do?)
Ayo, Maddie, what it do? (Hah)

[Chorus]
How can I lose if I'm already chose? Like
If she feelin' hot, then I make that bitch froze
And I get a bitch tight every time that I post, damn
The party not lit, then I'd rather not go (Why would I go?)
If she feelin' hot, then I make that bitch froze
And I get a bitch tight every time that I post, damn
""")
         
totalSongs = song1 + song2 + song3

model_engine = "text-davinci-002"
response = openai.Completion.create(
  engine=model_engine,
  prompt=totalSongs,
  max_tokens=1024,
  n=1,
  stop=None,
  temperature=0.7,
)

lyrics = response.choices[0].text.strip()
print(lyrics)
