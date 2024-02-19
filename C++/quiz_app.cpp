#include <iostream>
#include <string>

using namespace std;

const int NUM_QUESTIONS = 5;
struct Question
{
  string questionText;
  string options[4];
  char correctAnswer;
};

void displayQuestion(const Question &q)
{
  cout << q.questionText << endl;
  for (const string &option : q.options)
  {
    cout << option << endl;
  }
}

bool checkAnswer(const Question &q, char answer)
{
  return toupper(answer) == q.correctAnswer;
}

void displayScore(const int score, const int totalQuestions)
{
  cout << "Your final score is " << score << " out of " << totalQuestions << "." << endl;
}

int main()
{
  Question questions[NUM_QUESTIONS] = {
      {"What is the capital of France?", {"A) London", "B) Berlin", "C) Paris", "D) Madrid"}, 'C'},
      {"Who wrote 'To Kill a Mockingbird'?", {"A) F. Scott Fitzgerald", "B) Harper Lee", "C) Ernest Hemingway", "D) Jane Austen"}, 'B'},
      {"What is the currency of Japan?", {"A) Yen", "B) Dollar", "C) Euro", "D) Peso"}, 'A'},
      {"What is the largest ocean on Earth?", {"A) Atlantic Ocean", "B) Pacific Ocean", "C) Indian Ocean", "D) Arctic Ocean"}, 'B'},
      {"What is the smallest country in the world?", {"A) Monaco", "B) Vatican City", "C) London", "D) San Marino"}, 'B'}};

  int score = 0;
  char userAnswer;
  bool correct;

  for (int i = 0; i < NUM_QUESTIONS; ++i)
  {
    displayQuestion(questions[i]);
    cout << "Enter your answer (A, B, C, or D): ";
    cin >> userAnswer;
    correct = checkAnswer(questions[i], userAnswer);

    if (correct)
    {
      cout << "Correct!" << endl;
      score++;
    }
    else
    {
      cout << "Incorrect. The correct answer is " << questions[i].correctAnswer << "." << endl;
    }
    cout << endl
         << endl
         << endl
         << endl;
  }

  displayScore(score, NUM_QUESTIONS);
  return 0;
}
