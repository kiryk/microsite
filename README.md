# `microsite`

Microsite is a minimal personal homepage server that lets you post articles divided into sections.

## Usage

For `microsite` sections are directories and articles are files, both of these are meant to be managed through FTP or directly.

### Sections

To add a new section to the website, create a folder in the server's main directory.

The name of the folder, will be the name of the section, `microsite` will display the names with capitalized first letters of each word in sections name.

Articles in each section are sorted by date of _modification_, starting with the newest one.

An example of server's main direcory might look like this:

	/articles
	/projects
	/about me
	index.html
	microsite

### Articles

To add an article, create an HTML file in a section directory.

Names of the files are displayed as article titles in section view, and `microsite` will hide the extension (if there is any).

The expected format of a file containing an article is HTML encoded in UTF-8 and without the `<body>` tag, but only with the very content of the article. This format corresponts to output of many preprocessing tools like Markdown.

### Styling

The style used by the server by default is just exemplary, to change it - modify the `template.html` file. This file serves both as a template for section pages and article pages.

The CSS styling is embedded in the file, and it's rather intended to stay like that for both efficiency and just simplicity.

## To be done

I want to add some support for modules, like widgets etc., so the program could be easily extended.

I also plan to add built-in support for Markdown or its subset, if I ever come up with a solution minimalist enough.
