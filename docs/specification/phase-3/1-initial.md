Suites (1)
1 failed
Tests (17)
16 passed
1 failed

/home/node/dist/auto-test-checker/dist/tests/jest/back/back-v3-27-homework-01/quiz-game-top-players-flow.spec.js

Homework 27 > Top players
DELETE -> "/testing/all-data": should remove all data; status 204;
passed
0.129s

Homework 27 > Top players
POST -> "/sa/quiz/questions", PUT -> "/sa/quiz/questions/:questionId/publish": should create and publish several questions; status 201; content: created question;
passed
1.291s

Homework 27 > Top players
POST -> "/sa/users", "/auth/login": should create and login 5 users; status 201; content: created users;
passed
6.956s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user2, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.494s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user2, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.459s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user1, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.46s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user1, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.398s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user3, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; draw with 2 scores; ; status 200;
passed
1.471s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user4, then: add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; secondPlayer should win, scores: 5 - 0; no one got an extra point; ; status 200;
passed
1.437s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user4, connect to the game by user1, then: add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; secondPlayer should win, scores: 5 - 0; no one got an extra point; ; status 200;
passed
1.429s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user4, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.417s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user3, connect to the game by user4, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.409s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user5, then: add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; secondPlayer should win, scores: 5 - 0; no one got an extra point; ; status 200;
passed
1.423s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user5, then: add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; secondPlayer should win, scores: 5 - 0; no one got an extra point; ; status 200;
passed
1.585s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user5, connect to the game by user3, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.417s

Homework 27 > Top players
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user3, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.47s

Homework 27 > Top players
GET -> "/pair-game-quiz/users/top": expected statistics: user1: {"gamesCount":9,"winsCount":4,"lossesCount":4,"drawsCount":1,"sumScore":20,"avgScores":2.22}; user2: {"gamesCount":6,"winsCount":3,"lossesCount":3,"drawsCount":0,"sumScore":13,"avgScores":2.17}; user3: {"gamesCount":4,"winsCount":1,"lossesCount":2,"drawsCount":1,"sumScore":9,"avgScores":2.25}; user4: {"gamesCount":4,"winsCount":1,"lossesCount":3,"drawsCount":0,"sumScore":9,"avgScores":2.25}; user5: {"gamesCount":3,"winsCount":3,"lossesCount":0,"drawsCount":0,"sumScore":13,"avgScores":4.33}; status 200; content: list of top players;
failed
0.169s

> Homework 27  Top players  GET -> "/pair-game-quiz/users/top": expected statistics:
user1: {"gamesCount":9,"winsCount":4,"lossesCount":4,"drawsCount":1,"sumScore":20,"avgScores":2.22};
user2: {"gamesCount":6,"winsCount":3,"lossesCount":3,"drawsCount":0,"sumScore":13,"avgScores":2.17};
user3: {"gamesCount":4,"winsCount":1,"lossesCount":2,"drawsCount":1,"sumScore":9,"avgScores":2.25};
user4: {"gamesCount":4,"winsCount":1,"lossesCount":3,"drawsCount":0,"sumScore":9,"avgScores":2.25};
user5: {"gamesCount":3,"winsCount":3,"lossesCount":0,"drawsCount":0,"sumScore":13,"avgScores":4.33}; status 200; content: list of top players;

    expect(received).toBe(expected) // Object.is equality

    Expected: 200
    Received: 401

      102 |     }
      103 |
    > 104 |     expect(status).toBe(expectedStatusCode);
          |                    ^
      105 |
      106 |     if (expectedData) {
      107 |       expect(data).toBeEqualWithQueryParams(expectedData, queryParams, withDiffPrint);

      at performQueryParamsChecker (src/tests/jest/back/testHelpers/performCheckers.ts:104:20)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-top-players-describe.ts:85:7)
