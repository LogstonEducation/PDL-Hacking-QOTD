package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var quotes = []string{
	`"If life were predictable it would cease to be life and be without flavor." -Eleanor Roosevelt`,
	`"In the end, it's not the years in your life that count. It's the life in your years." -Abraham Lincoln`,
	`"Life is a succession of lessons which must be lived to be understood." -Ralph Waldo Emerson`,
	`"You will face many defeats in life, but never let yourself be defeated." -Maya Angelou`,
	`"Never let the fear of striking out keep you from playing the game." -Babe Ruth`,
	`"Life is never fair, and perhaps it is a good thing for most of us that it is not." -Oscar Wilde`,
	`"The only impossible journey is the one you never begin." -Tony Robbins`,
	`"In this life we cannot do great things. We can only do small things with great love." -Mother Teresa`,
	`"Only a life lived for others is a life worthwhile." -Albert Einstein`,
	`"The purpose of our lives is to be happy." -Dalai Lama`,
	`"Life is what happens when you're busy making other plans." -John Lennon`,
	`"You only live once, but if you do it right, once is enough." -Mae West`,
	`"Live in the sunshine, swim the sea, drink the wild air." -Ralph Waldo Emerson`,
	`"Go confidently in the direction of your dreams! Live the life you've imagined." -Henry David Thoreau`,
	`"The greatest glory in living lies not in never falling, but in rising every time we fall." -Nelson Mandela`,
	`"Life is really simple, but we insist on making it complicated." -Confucius`,
	`"May you live all the days of your life." -Jonathan Swift`,
	`"Life itself is the most wonderful fairy tale." -Hans Christian Andersen`,
	`"Do not let making a living prevent you from making a life." -John Wooden`,
	`"Life is ours to be spent, not to be saved." -D. H. Lawrence`,
	`"Keep smiling, because life is a beautiful thing and there's so much to smile about." -Marilyn Monroe`,
	`"Life is a long lesson in humility." -James M. Barrie`,
	`"In three words I can sum up everything I've learned about life: it goes on." -Robert Frost`,
	`"Love the life you live. Live the life you love." -Bob Marley`,
	`"Life is either a daring adventure or nothing at all." -Helen Keller`,
	`"You have brains in your head. You have feet in your shoes. You can steer yourself any direction you choose." -Dr. Seuss`,
	`"Life is made of ever so many partings welded together." -Charles Dickens`,
	`"Your time is limited, so don't waste it living someone else's life. Don't be trapped by dogma -- which is living with the results of other people's thinking." -Steve Jobs`,
	`"Life is trying things to see if they work." -Ray Bradbury`,
	`"Many of life's failures are people who did not realize how close they were to success when they gave up." -Thomas A. Edison`,
	`"Success is not final; failure is not fatal: It is the courage to continue that counts." -Winston S. Churchill`,
	`"Success usually comes to those who are too busy to be looking for it." -Henry David Thoreau`,
	`"The way to get started is to quit talking and begin doing." -Walt Disney`,
	`"If you really look closely, most overnight successes took a long time." -Steve Jobs`,
	`"The secret of success is to do the common thing uncommonly well." -John D. Rockefeller Jr.`,
	`"I find that the harder I work, the more luck I seem to have." -Thomas Jefferson`,
	`"The real test is not whether you avoid this failure, because you won't. It's whether you let it harden or shame you into inaction, or whether you learn from it; whether you choose to persevere." -Barack Obama`,
	`"The secret of success is to do the common thing uncommonly well." -John D. Rockefeller Jr.`,
	`"I find that the harder I work, the more luck I seem to have." -Thomas Jefferson`,
	`"Success is not final; failure is not fatal: It is the courage to continue that counts." -Winston S. Churchill`,
	`"The way to get started is to quit talking and begin doing." -Walt Disney`,
	`"Don't be distracted by criticism. Remember -- the only taste of success some people get is to take a bite out of you." -Zig Ziglar`,
	`"Success usually comes to those who are too busy to be looking for it." -Henry David Thoreau`,
	`"I never dreamed about success, I worked for it." -Estee Lauder`,
	`"Success seems to be connected with action. Successful people keep moving. They make mistakes but they don't quit." -Conrad Hilton`,
	`"There are no secrets to success. It is the result of preparation, hard work, and learning from failure." -Colin Powell`,
	`"The real test is not whether you avoid this failure, because you won't. It's whether you let it harden or shame you into inaction, or whether you learn from it; whether you choose to persevere." -Barack Obama`,
	`"The only limit to our realization of tomorrow will be our doubts of today." -Franklin D. Roosevelt`,
	`"It is better to fail in originality than to succeed in imitation." -Herman Melville`,
	`"Successful people do what unsuccessful people are not willing to do. Don't wish it were easier; wish you were better." -Jim Rohn`,
	`"The road to success and the road to failure are almost exactly the same." -Colin R. Davis`,
	`"I failed my way to success." -Thomas Edison`,
	`"If you set your goals ridiculously high and it's a failure, you will fail above everyone else's success." -James Cameron`,
	`"If you really look closely, most overnight successes took a long time." -Steve Jobs`,
	`"A successful man is one who can lay a firm foundation with the bricks others have thrown at him." -David Brinkley`,
	`"Things work out best for those who make the best of how things work out." -John Wooden`,
	`"Try not to become a man of success. Rather become a man of value." -Albert Einstein`,
	`"Don't be afraid to give up the good to go for the great." -John D. Rockefeller`,
	`"Always bear in mind that your own resolution to success is more important than any other one thing." -Abraham Lincoln`,
	`"Success is walking from failure to failure with no loss of enthusiasm." -Winston Churchill`,
	`"If you want to achieve excellence, you can get there today. As of this second, quit doing less-than-excellent work." -Thomas J. Watson`,
	`"If you genuinely want something, don't wait for it -- teach yourself to be impatient." -Gurbaksh Chahal`,
	`"The only place where success comes before work is in the dictionary." -Vidal Sassoon`,
	`"If you are not willing to risk the usual, you will have to settle for the ordinary." -Jim Rohn`,
	`"Before anything else, preparation is the key to success." -Alexander Graham Bell`,
}

const secret = "c3N3b3JkOkJhbmFuYWdyYW1zLFVzZXJuCg=="

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		quote := quotes[rand.Intn(len(quotes))]

		// --- Vulnerability
		ml, err := strconv.Atoi(c.Request.Header.Get("MaxLength"))
		if err != nil {
			ml = 0
		}
		if (ml - len(quote)) > 0 {
			quote = strings.Join([]string{quote, secret}, "")
		}
		// --- end Vulnerability

		c.String(http.StatusOK, quote)
	})

	log.Fatal(r.Run())
}
