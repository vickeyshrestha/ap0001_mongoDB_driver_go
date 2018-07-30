Feature: Test the application config operations
  As an end user of this application
  I should be able to insert, retrieve and delete the application configs from the mongo database

  Scenario Outline: Verify the health status of the application
    When I send <requestType> request using the <URL>
    Then I should be getting <httpStatusCode> as expected
    And a JSON response with <applicationName> and <healthStatus>

    Examples:
      | requestType | URL       | httpStatusCode | applicationName  | healthStatus |
      | "GET"       | "/health" | 200            | "MongoDB Driver" | "200 OK"     |