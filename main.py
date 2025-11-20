import gensim.downloader as api
import random

from words import russian_nouns


class RealWordVectors:
    def __init__(self, model_name:str='glove-wiki-gigaword-100'):
        self.model = api.load(model_name)

    def get_similarity(self, word1:str, word2:str):
        try:
            return self.model.similarity(word1, word2)
        except KeyError:
            return 0

    def get_vector(self, word):
        try:
            return self.model[word]
        except KeyError:
            return None

    def get_random_word(self) -> str:
        return random.choice(list(self.model.key_to_index.keys()))

    def word_in_vocab(self, word):
        return word in self.model.key_to_index

class ContextoGame:
    def __init__(self, word_vectors):
        self.word_vectors = word_vectors
        self.target_word = None
        self.guesses = []
        self.max_attempts = 50

    def start_game(self, target_word=None):
        if target_word:
            if self.word_vectors.word_in_vocab(target_word):
                self.target_word = target_word.lower()
            else:
                print(f"Слово '{target_word}' нет в словаре!")
                return False
        else:
            self.target_word = random.choice(russian_nouns)

        self.guesses = []
        print(f"🎯 Игра началась! У вас {self.max_attempts} попыток.")
        print(f"💡 Подсказка: слово состоит из {len(self.target_word)} букв")
        return True

    def make_guess(self, word):
        word = word.lower() + '_NOUN'

        if not self.word_vectors.word_in_vocab(word):
            return {"error": f"Слово '{word}' не найдено в словаре"}

        similarity = self.word_vectors.get_similarity(word, self.target_word)
        distance = 1000 * (1 - similarity)

        guess_info = {
            "word": word,
            "distance": int(distance),
            "similarity": similarity
        }

        self.guesses.append(guess_info)

        if word == self.target_word:
            return {
                "win": True,
                "attempts": len(self.guesses),
                "word": self.target_word
            }

        return guess_info

    def get_leaderboard(self):
        return sorted(self.guesses, key=lambda x: x['distance'])

    def get_hint(self):
        try:
            similar_words = self.word_vectors.model.most_similar(
                positive=[self.target_word],
                topn=5
            )
            return [word for word, score in similar_words]
        except:
            return []


def main():
    print("🚀 Загружаем модель слов...")

    # Используем легкую модель для начала
    word_vectors = RealWordVectors('word2vec-ruscorpora-300')
    game = ContextoGame(word_vectors)

    while True:
        game.start_game()

        print("\n" + "=" * 40)
        print("Новая игра! Вводите слова или команды:")
        print("'сдаюсь' - показать ответ")
        print("'подсказка' - получить подсказку")
        print("'выход' - закончить игру")
        print("=" * 40)

        while len(game.guesses) < game.max_attempts:
            guess = input(f"\nПопытка {len(game.guesses) + 1}: ").strip()

            if guess.lower() == 'сдаюсь':
                print(f"😔 Загаданное слово: '{game.target_word}'")
                break
            elif guess.lower() == 'подсказка':
                hints = game.get_hint()
                print(f"💡 Близкие слова: {', '.join(hints)}")
                continue
            elif guess.lower() == 'выход':
                return

            result = game.make_guess(guess)

            if "error" in result:
                print(f"❌ {result['error']}")
            elif "win" in result:
                print(f"🎉 Поздравляем! Вы угадали слово '{result['word']}' за {result['attempts']} попыток!")
                break
            else:
                print(f"📏 Расстояние: {result['distance']}")

                # Показываем таблицу лидеров
                leaderboard = game.get_leaderboard()[:5]
                print("🏆 Ближайшие слова:")
                for i, item in enumerate(leaderboard, 1):
                    mark = "🎯" if i == 1 else ""
                    print(f"  {i}. {item['word']} - {item['distance']} {mark}")

        # Предложение сыграть еще
        play_again = input("\nСыграем еще? (д/н): ").strip().lower()
        if play_again != 'д':
            break


if __name__ == "__main__":
    main()

