package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var prompts = []string{
	"In a quiet village nestled between rolling hills, the weather is a tranquil 24 degrees Celsius. The sky is painted in hues of orange and pink as the sun sets, casting a warm glow across the landscape. Residents enjoy the pleasant evening breeze, a perfect setting for relaxation and reflection.",
	"Embark on a culinary journey with a delightful recipe for Mediterranean Stuffed Peppers. Packed with vibrant vegetables, fragrant herbs, and seasoned rice, these peppers are a taste sensation. Whether for a family dinner or a festive gathering, this dish is sure to impress.",
	"Step into the realm of classic cinema with 'Casablanca,' a timeless masterpiece released in 1942. Set against the backdrop of World War II, the film unfolds a tale of love, sacrifice, and moral dilemmas. Its enduring charm continues to captivate audiences across generations.",
	"Journey into the world of languages as you translate the phrase 'Thank you for your kindness' from English to Japanese. '親切にありがとう' gracefully expresses gratitude in Japanese, showcasing the beauty of cultural diversity and linguistic nuances.",
	"Navigate the dynamic stock market with insights into TechGrowth Inc. The current stock price stands at $120.85, reflecting a 2.5% increase. Analysts predict a positive trajectory, making it an opportune time for investors to monitor this emerging tech giant.",
	"Delve into the latest breakthroughs in artificial intelligence, where machine learning algorithms redefine industries. TechInnovate News explores how AI applications are transforming healthcare, finance, and logistics, paving the way for a future shaped by intelligent technologies.",
	"Experience the thrill of a nail-biting basketball game between the Bulls and the Celtics. In a high-scoring match, the Bulls emerged victorious with a final score of 118, securing a well-deserved win. Sports enthusiasts can relish the intense moments that unfolded on the court.",
	"Lighten the mood with a chuckle-worthy dad joke: 'Why did the computer keep its drink on the windowsill? Because it wanted a byte with a view!' Embrace the whimsy and humor in everyday situations, making the journey through tech and life more enjoyable.",
	"Under the vast night sky, a telescope reveals the wonders of distant galaxies. Stars twinkle in a cosmic ballet, and nebulae paint the canvas of space. The telescope, a gateway to the universe, invites observers to marvel at the beauty and mysteries of the celestial realm.",
	"Embark on a coding adventure, exploring the versatility of Python. From web development to data analysis, Python's simplicity and power make it a language of choice for programmers. Unleash your creativity and build innovative solutions with this versatile scripting language.",
	"Ignite your artistic passion with a blank canvas and vibrant colors. Whether it's a serene landscape or an abstract masterpiece, the world of art offers endless possibilities. Channel your emotions onto the canvas and let your creativity flow in every brushstroke.",
	"Explore the enchanting world of underwater ecosystems with scuba diving. Witness the kaleidoscope of marine life, from colorful coral reefs to majestic sea turtles. Dive into the depths and discover the beauty and biodiversity hidden beneath the ocean's surface.",
	"Savor the rich aroma of freshly brewed coffee as you craft your perfect cup. Whether you prefer a velvety espresso or a frothy cappuccino, the ritual of brewing coffee is an art. Take a moment to appreciate the nuanced flavors and the comforting warmth in every sip.",
	"Dive into the pages of a gripping mystery novel, where suspense lurks around every corner. The plot thickens as detectives unravel clues, leading to an unexpected twist. Immerse yourself in the narrative, and let the suspenseful tale transport you to another world.",
	"Stay active and embrace the joy of movement with a dance workout. Feel the rhythm, let loose, and enjoy the positive effects on both body and mind. Dancing is not just exercise; it's a celebration of self-expression and a pathway to physical well-being.",
	"Immerse yourself in the captivating melodies of classical music. From the symphonies of Beethoven to the operas of Mozart, classical compositions stand the test of time. Let the harmonies and intricate arrangements transport you to a world of timeless musical elegance.",
	"Appreciate the intricate details of architecture as you explore historic landmarks. From Gothic cathedrals to modern skyscrapers, architecture is a testament to human creativity and ingenuity. Take a stroll through architectural wonders and witness the beauty of design.",
	"Capture the essence of a tropical paradise with a refreshing coconut smoothie. Blending coconut water, pineapple, and a hint of mint creates a tropical symphony of flavors. Indulge in this revitalizing drink, transporting your taste buds to a sun-kissed island getaway.",
	"Unleash your inner storyteller with the magic of words. Whether crafting a tale of fantasy or penning a thought-provoking poem, writing is a journey of self-expression. Navigate the realms of imagination and bring your narratives to life with the power of storytelling.",
	"Embark on a virtual tour of renowned art museums, exploring masterpieces from different periods. From Renaissance classics to contemporary wonders, the world of art awaits. Wander through the digital galleries and appreciate the beauty of artistic expressions spanning centuries.",
}

func main() {
	app := fiber.New()

	app.Get("/promt", randomPrompt)

	log.Fatal(app.Listen(":3030"))
}

func randomPrompt(ctx *fiber.Ctx) error {
	index := rand.Intn(20)
	prompt := prompts[index]
	return ctx.Status(http.StatusOK).SendString(prompt)
}
