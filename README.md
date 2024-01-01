# [Athena](https://github.com/kyallanum/athena)
[![License](https://img.shields.io/badge/License-Apache_2.0-green.svg)](https://opensource.org/licenses/Apache-2.0) 
[![Build](https://img.shields.io/github/actions/workflow/status/kyallanum/athena/go.yml)](https://github.com/kyallanum/athena/actions)  

A universal log parser. Weed out the unimportant information.

## Requirements
- **Windows x86/amd64**
- **Mac OS**
- **Linux x86/amd64**

> [!NOTE]  
> Currently for Mac OS, the binary must be marked as trusted in order to execute. This binary has not yet been signed.

## Purpose
Athena is a product that allows users to look for specific entries in log files. With the use of regular expressions, it both prints out log lines that it has found, along with adding critical information in memory, to then be used later in a summary. This is primarily meant for users that would like a summary of what occurred in a large log file, and to discern the important pieces of a log.

> [!NOTE]  
> Athena is not a ML model and does not detect critical pieces of information in a log file by itself. But rather makes use of a configuration file to look for the information that it needs.

## Execution
Athena has three main phases of execution:
1. File ingest (configuration and log file)
1. Log File Resolution
1. Summary Output

To run Athena, you can use either environment variables, or Command Line Flags. As a rule, environment variables take precedence over command line flags. Please see [CLI Flags](#cli-flags) for more info.

## CLI Flags
Athena has a few different command-line flags/environment variables a user needs to be aware of:
| Flag            | Description                           | Required | Default | Environment Variable |
|-----------------|---------------------------------------|----------|---------|----------------------|
| -c/--config     | Configuration file for Athena to use. |   true   | N/A     | ATHENA_CONFIG_FILE   |
| -l/--log-file   | The log file Athena should process.   |   true   | N/A     | ATHENA_LOG_FILE      |
| -o/--log-output | The log file Athena should output to. |   false  | N/A     | ATHENA_LOG_OUTPUT    |

## Configuration File
An athena configuration file is a JSON file that describes how Athena should process a log file. It makes use of a hierarchy of objects for instructions during execution. This can either be from a URL or a local file.

### The Log File
The top level of the configuration file is at the Log file level. This holds two different pieces of information. The name and the rules for the log file.

### Rules
A rule should be considered certain information that Athena must follow to provide proper information to the user. A rule should focus on one specific type of information to be extracted. It includes the following parameters:  
1. Name - The name and/or purpose of this rule
1. PrintLog - A boolean value describing whether a log line should be printed to the console when it finds relevant information.
1. [Search Terms](#search-terms)
1. [Summary](#summary)

### Search Terms
Search terms are a list of strings, in regular expression format that can be used to match to a line in a log file. Named groups using the format ``(?<group_name>)`` can be used to store information in memory. This can later be accessed using the format: ``{{group_name}}``.

**Guidelines for search terms**:
1. You cannot reference a group that has not been defined in a previous search term.
1. The regular expression must be valid, Athena does not support [lookaheads or lookbehinds](https://www.regular-expressions.info/lookaround.html/).

### Summary
The summary is a list of strings to be printed out in the end. The summary itself is printed out in the format:
```
--------------- <Log File Name> Log File Summary ---------------
Rule <Rule Name>:
<Output Lines>
```

**Guidelines for a Rule Summary**:
1. You can reference the named groups previously resolved in the search phase. This is done with the format:  
 ``{{<operation>(<group_name>)}}``
1. Operations manipulate the information stored in some way for printing out in the Summary.
1. This syntax is very limited, and currently only supports one type of operation per line (you cannot combine a count and a print operation on the the same line.)

**Current Supported Operations**:  
1. Count - Counts the number of times this ``<group_name>`` was extracted during the search phase.
1. Print - Prints out a line for every instance of ``<group_name>`` that was extracted. If there are two "Print" operations, then the second one's data will match to the first one in each line.

## Log Output
Athena's logging is designed to output to the console with normal output, while outputting to a file with timestamps, log levels, and more. If you would like all information including debug logs, then set the ``ATHENA_LOG_OUTPUT`` environment variable or the ``-o/--log-output`` command-line flag.

---
##### Licensed under [Apache 2.0 License](https://opensource.org/license/apache-2-0/) (c) Kyal Lanum
