> Homework 26  My games  GET -> "/pair-game-quiz/pairs/my": create, connect to game 3 times and add all answers sequantly. Then create, connect 4th game (not finished). Then get 'myGames' sorted by 'status'; status 200; content: list of current user games (finished and current);  used additional methods: POST -> /pair-game-quiz/pairs/connection, POST -> /pair-game-quiz/pairs/my-current/answers;

    expect(received).toBe(expected) // Object.is equality

    Expected: 200
    Received: 400

      102 |     }
      103 |
    > 104 |     expect(status).toBe(expectedStatusCode);
          |                    ^
      105 |
      106 |     if (expectedData) {
      107 |       expect(data).toBeEqualWithQueryParams(expectedData, queryParams, withDiffPrint);

      at performQueryParamsChecker (src/tests/jest/back/testHelpers/performCheckers.ts:104:20)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-myGames-describe.ts:106:7)

> Homework 26  My statistic  GET -> "/pair-game-quiz/users/my-statistic": should return status 200; content: current user's games statistic;



    Expected: success response

    Received: Request failed with status code 404

    Config:
     url: pair-game-quiz/users/my-statistic
     method: get
     response status: 404
     request body: undefined
     response data: {"errorsMessages":[{"message":"Not Found","field":""}]}

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-myStatistic-describe.ts:44:35
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-myStatistic-describe.ts:43:40)
