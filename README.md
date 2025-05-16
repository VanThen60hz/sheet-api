# personel-api

A Golang API for users to programmatically manage Google Sheets data.

## How to Use

Step 1:
Navigate to root folder

Step 2:

    Run:
        $env:testing = "false"; go run ./cmd/main.go

    Test:
        $env:testing = "true"; go test ./... -coverprofile=coverage

    Generate coverage HTML:
        go tool cover -html=coverage -o coverage.html

## Running the Project

### Prerequisites

-   Go 1.16 or later installed on your machine
-   Git for version control
-   Access to Google Sheets API (see credentials section below)

### Step-by-step Instructions

1. Clone the repository:

    ```
    git clone https://gitlab.com/mksgroup/intern/personinfo/personel-api.git
    cd personel-api
    ```

2. Install dependencies:

    ```
    go mod download
    ```

3. Set up credentials:

    - Make sure you have `credentials.json` in the project root directory
    - Generate `token.json` as described in the Admin section below

4. Run the project:

    - For development mode:
        ```
        $env:testing = "false"; go run ./cmd/main.go
        ```
    - The API server will start at http://localhost:8080 (default port)

5. Testing:
    - Run tests:
        ```
        $env:testing = "true"; go test ./... -coverprofile=coverage
        ```
    - Generate and view coverage report:
        ```
        go tool cover -html=coverage -o coverage.html
        ```

### Build and Deploy

To build the project for production:

```
go build -o personnel-api ./cmd/main.go
```

Run the built binary:

```
./personnel-api
```

## Function Description

## GET

### GetAll [get]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

    Des:
        Get all data from a spreadsheet with spreadsheetID in json format.

### GetSheetData [get]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

    Des:
        Get all data from a sheet with sheetName.

### GetByColumn [get]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - columnName (required)
            Type: String
            Description: Name of the column you want to read from (The first row of the column should be the column name).

    Des:
        Get all data from a column with columnName.

### GetByFilter [get]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - columnName (required)
            Type: String
            Description: Name of the column you want to read from (The first row of the column should be the column name).

        - operator (required)
            Type: String
            Description: int and float accept ">", "<", "=". String accepts "contain".

        - value (required)
            Type: String
            Description: Value to compare against.

    Des:
        Get all data from a column that passes the filter.

## Create

### CreateData [post]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - rows (required)
            Type: [][]interface{}
            Description: Rows of data to be appended.

    Des:
        Append data to a specific sheet.

## Delete

### DeleteDataRow [delete]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - range (required)
            Type: []interface{}
            Description: Indexes of rows to be deleted.

    Des:
        Delete data from specific rows.

### DeleteDataCell [delete]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - range (required)
            Type: [][]interface{}
            Description: Coordinates of the cells to be deleted.

    Des:
        Delete data from specific cells.

## Update

### UpdateDataRow [put]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - rows (required)
            Type: [][]interface{}
            Description: New data of rows to be updated.

        - range (required)
            Type: []interface{}
            Description: Indexes of rows to be updated.

    Des:
        Update data of specific rows.

### UpdateDataCell [put]

    Param:
        - spreadsheetID (required)
            Type: String
            Description: The unique identifier (ID) associated with the target spreadsheet.

        - sheetName (required)
            Type: String
            Description: Name of the data sheet you want to read from.

        - cells (required)
            Type: []interface{}
            Description: New data of cells to be updated.

        - range (required)
            Type: [][]interface{}
            Description: Coordinates of cells to be updated.

    Des:
        Update data of specific cells.

## For Admin

### Activating Google Sheets API:

For production, please replace both token.json and credential.json files with your own files.

Instruction: https://developers.google.com/sheets/api/quickstart/go

### Obtaining New Credentials and Token Files:

#### Getting a new credentials.json file:

1. Go to the [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Enable the Google Sheets API for your project
4. Go to "Credentials" in the left sidebar
5. Click "Create Credentials" and select "OAuth client ID"
6. Select "Desktop app" as the application type
7. Name your OAuth client and click "Create"
8. Download the credentials by clicking the download icon (JSON)
9. Rename the downloaded file to `credentials.json` and place it in the project root directory

#### Getting a new token.json file:

1. Delete your existing token.json file (if any)
2. Run the application once with your new credentials.json
3. The application will prompt you to authorize access:
    - A browser window will open automatically
    - Log in with your Google account
    - Grant the requested permissions
4. After authorization, a new token.json file will be automatically generated in the project root

Note: The token.json file contains access tokens and should be kept secure and not committed to version control.

### CASBIN:

Please replace admin_key in main.go with your personal key.

## TODO:

I have set up a model for Casbin and created a policy.csv file. You can use it to finish authorization/Auth.go
to finish setting up authorization.

Implement Swagger for debugging this API.

Implement CreateSheet function which allows users to create a new Sheet programmatically.

Implement CreateColumnName function which allows users to set names for new columns.

## Guide for New Developers

This section provides comprehensive guidance for new developers joining the project.

### Project Structure Overview

```
personnel-api/
├── cmd/                  # Entry point for the application
│   ├── main.go           # Main application file
│   └── docs/             # API documentation
├── pkg/                  # Reusable packages
│   ├── api/              # API implementation
│   │   ├── create/       # Create operations
│   │   ├── read/         # Read operations
│   │   ├── update/       # Update operations
│   │   └── delete/       # Delete operations
│   ├── authorization/    # Authentication and authorization
│   └── svc/              # Core services
├── credentials.json      # Google API credentials
├── token.json            # Google API access token
├── model.conf            # CASBIN model configuration
├── policy.csv            # CASBIN policy definitions
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── .gitlab-ci.yml        # GitLab CI/CD configuration
└── README.md             # Project documentation
```

Note: The project currently doesn't use an internal folder structure. The code organization is primarily based on the pkg directory.

### Setup Development Environment

1. Follow the installation steps in the "Running the Project" section above.
2. Make sure you have a code editor with Go support (VS Code with Go extension recommended).
3. Install required development tools:
    ```
    go install golang.org/x/tools/cmd/goimports@latest
    go install golang.org/x/lint/golint@latest
    ```

### Understanding the Codebase

1. Start by examining `cmd/main.go` to understand the application initialization and server setup.
2. The API is organized in the `pkg/api` directory with separate modules for different operations:
    - `pkg/api/read/` - Contains all read operations (GET endpoints)
    - `pkg/api/create/` - Contains all create operations (POST endpoints)
    - `pkg/api/update/` - Contains all update operations (PUT endpoints)
    - `pkg/api/delete/` - Contains all delete operations (DELETE endpoints)
3. Authorization is handled in the `pkg/authorization/` directory
4. Core services and Google Sheets API integration are in the `pkg/svc/` directory
5. The data flow follows this pattern:
    - Request → API Handler (in pkg/api/\*) → Service (in pkg/svc) → Google Sheets API

### Adding a New Feature

Follow these steps to add a new feature:

1. **Understand the Requirements**

    - Clearly define what the new feature should do
    - Identify which parts of the codebase will need to be modified

2. **Create a New Branch**

    ```
    git checkout -b feature/your-feature-name
    ```

3. **Implement the Feature**

    - For a new API endpoint:

        1. Create a new handler function in the appropriate file in `internal/handlers/`
        2. Add the business logic in `internal/service/`
        3. If needed, add data access functions in `internal/repository/`
        4. Register the new route in the router setup (usually in `cmd/main.go` or a separate router file)

    - For extending existing functionality:
        1. Identify the relevant handlers, services, and repositories
        2. Add or modify the code as needed

4. **Implement the Feature**

    - For a new API endpoint:

        1. Identify which operation type your endpoint belongs to (read, create, update, or delete)
        2. Create a new file or modify an existing file in the appropriate directory under `pkg/api/`
        3. If needed, add new service functions in `pkg/svc/`
        4. Register the new route in `cmd/main.go`

    - For extending existing functionality:
        1. Identify the relevant files in `pkg/api/` and `pkg/svc/`
        2. Add or modify the code as needed
        3. Ensure proper error handling and response formatting

5. **Write Tests**

    - Create unit tests for your new code
    - Update existing tests if you modified existing code

    ```
    $env:testing = "true"; go test ./path/to/your/package -v
    ```

6. **Test Your Feature**

    - Run the application locally
    - Test the new feature using appropriate tools (curl, Postman, etc.)

7. **Submit Your Changes**

    ```
    git add .
    git commit -m "Add feature: your feature description"
    git push origin feature/your-feature-name
    ```

    - Create a merge request on GitLab

    ### Common Challenges and Solutions

8. **Google Sheets API Authentication Issues**

    - Ensure your `credentials.json` is correctly set up as described in the Admin section
    - If authentication fails, delete `token.json` and let the application regenerate it

9. **Understanding the Data Flow**

    - The application follows a standard pattern: HTTP Request → Handler → Service → Repository → Google Sheets API
    - Each layer has a specific responsibility, maintaining separation of concerns

10. **Error Handling Best Practices**
    - Use appropriate HTTP status codes in handlers
    - Log errors with sufficient context
    - Return structured error responses to clients

### Debugging Tips

1. Add logging statements using the standard Go log package or a custom logger
2. Use the Go debugger in your IDE (e.g., VS Code's Go debugger)
3. For API testing, use tools like Postman or curl to send requests and inspect responses

### Code Style and Standards

1. Follow standard Go coding conventions
2. Use goimports to organize imports
3. Run golint before committing code
4. Write clear, concise comments for functions and complex logic

## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

-   [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
-   [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://gitlab.com/mksgroup/intern/personinfo/personel-api.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

-   [ ] [Set up project integrations](https://gitlab.com/mksgroup/intern/personinfo/personel-api/-/settings/integrations)

## Collaborate with your team

-   [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
-   [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
-   [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
-   [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
-   [ ] [Automatically merge when pipeline succeeds](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

-   [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
-   [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing(SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
-   [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
-   [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
-   [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

---

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README

Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name

Choose a self-explaining name for your project.

## Description

Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges

On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals

Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation

Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage

Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support

Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap

If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing

State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment

Show your appreciation to those who have contributed to the project.

## License

For open source projects, say how it is licensed.

## Project status

If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
#   s h e e t - a p i  
 #   s h e e t - a p i  
 #   s h e e t - a p i  
 #   s h e e t - a p i  
 