# GoProj18
Instructions for Project 2018 for Graph Theory 

3rd Year Software Development - GMIT


## A Program in GO 
![](https://pbs.twimg.com/profile_images/867034392891457537/2_PUj0BC_400x400.jpg)



# Installing go
Please visit this link: https://golang.org/doc/install?download=go1.5.windows-amd64.msi2



# Instructions 
- Step 1: In cmd or terminal, git clone https://github.com/smoran98/GT18.git
- Step 2: Navigate to, GoProj18
- Step 3: go build main.go
- Step 4: main.exe or go run main.go


# Context
Real-world problem , e.g. industry. Expected to research and tackle the problem, use sources as it's been done in other other programming languages.



# Problem
Write a program in Go that can build a non-deterministic finite automaton (NFA) from a regular expression, and use the NFA to check if the regular expression matches any given string


*must write the program from scratch


A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “∗ ” have the special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1∗ means any number of 1’s. These special characters must be used in your submission

Other special characters you might consider allowing as input are brackets “()” which can be used for grouping, “+” which means “at least one of”, and “?” which means “zero or one of”. You might also decide to remove the concatenation character, so that 1.0 becomes 10, with the concatenation implicit.

You may initially restrict the non-special characters your program works with to 0 and 1, if you wish. However, you should at least attempt to expand these to all of the digits, and the characters a to z, and A to Z. You are expected to be able to break this project into a number of smaller tasks that are easier to solve, and to plug these together after they have been completed. 

# You might do that for this project as follows:

1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.

# Overall your program might have the following layout.
type nfa struct 
{


...


}


func regexcompile(r string) nfa {


...


return n


}


func (n nfa) regexmatch(n nfa, r sting) bool {


...


return ismatch


}


func main() {


n := regexcompile("01*0")


t := n.regexmatch("01110")


f := n.regexmatch("1000001")


}






# Requirements
GitHub repository containing a single Go program that can execute regular expressions against strings over the alphabet 0, 1 and with the special characters ., |, and ∗. 
README should clearly document how to compile, run and test your program and explain how your program works, and how you wrote it. Project should be well organised and contain detailed explanations. Architecture of the system to be well conceived, and examples of running the program be provided.



# Submissions
GitHub must be used to manage the development of the software. Your GitHub repo forms the main submission of the project and must submit the URL of your GitHub repository using the link on the course Moodle page before the deadline. You can do this at any time, as the last commit before the deadline will be used as your submission for this project.



# Marking scheme
25% Research Investigation of problem and possible solutions.


25% Development Clear architecture and well-written code.


25% Consistency Good planning and pragmatic attitude to work.


25% Documentation Detailed descriptions and explanations.



# REFERENCES
GMIT. Quality assurance framework


https://www.gmit.ie/general/quality-assurance-framework.


Google. The go programming language.


https://golang.org/.

