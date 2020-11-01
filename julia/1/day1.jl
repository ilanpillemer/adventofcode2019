{
 "cells": [
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "--- Day 1: The Tyranny of the Rocket Equation ---\n",
    "\n",
    "Santa has become stranded at the edge of the Solar System while delivering presents to other planets! To accurately calculate his position in space, safely align his warp drive, and return to Earth in time to save Christmas, he needs you to bring him measurements from fifty stars.\n",
    "\n",
    "Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!\n",
    "\n",
    "The Elves quickly load you into a spacecraft and prepare to launch.\n",
    "\n",
    "At the first Go / No Go poll, every Elf is Go until the Fuel Counter-Upper. They haven't determined the amount of fuel required yet.\n",
    "\n",
    "Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.\n",
    "\n",
    "For example:\n",
    "\n",
    "For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.\n",
    "For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.\n",
    "For a mass of 1969, the fuel required is 654.\n",
    "For a mass of 100756, the fuel required is 33583.\n",
    "The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.\n",
    "\n",
    "What is the sum of the fuel requirements for all of the modules on your spacecraft?"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 16,
   "metadata": {},
   "outputs": [],
   "source": [
    "fuel(m) = div(m,3) - 2;"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 15,
   "metadata": {},
   "outputs": [
    {
     "data": {
      "text/plain": [
       "true"
      ]
     },
     "execution_count": 15,
     "metadata": {},
     "output_type": "execute_result"
    }
   ],
   "source": [
    "[fuel(12), fuel(14), fuel(1969), fuel(100756)] == [2,2,654,33583]"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 24,
   "metadata": {},
   "outputs": [
    {
     "data": {
      "text/plain": [
       "3266516"
      ]
     },
     "execution_count": 24,
     "metadata": {},
     "output_type": "execute_result"
    }
   ],
   "source": [
    "using DelimitedFiles\n",
    "modules = readdlm(\"input.txt\", '\\t', Int, '\\n')\n",
    "sum(fuel.(modules))"
   ]
  },
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "--- Part Two ---\n",
    "\n",
    "During the second Go / No Go poll, the Elf in charge of the Rocket Equation Double-Checker stops the launch sequence. Apparently, you forgot to include additional fuel for the fuel you just added.\n",
    "\n",
    "Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2. However, that fuel also requires fuel, and that fuel requires fuel, and so on. Any mass that would require negative fuel should instead be treated as if it requires zero fuel; the remaining mass, if any, is instead handled by wishing really hard, which has no mass and is outside the scope of this calculation.\n",
    "\n",
    "So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative. For example:\n",
    "\n",
    "A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.\n",
    "At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.\n",
    "The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.\n",
    "What is the sum of the fuel requirements for all of the modules on your spacecraft when also taking into account the mass of the added fuel? (Calculate the fuel requirements for each module separately, then add them all up at the end.)"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 39,
   "metadata": {},
   "outputs": [],
   "source": [
    "function fuel(m) \n",
    "    total = 0\n",
    "    x = div(m,3) - 2\n",
    "    while x > 0\n",
    "        total = total + x\n",
    "        m = x\n",
    "        x = div(m,3) - 2\n",
    "    end  \n",
    "    total\n",
    "end;"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 45,
   "metadata": {},
   "outputs": [
    {
     "data": {
      "text/plain": [
       "true"
      ]
     },
     "execution_count": 45,
     "metadata": {},
     "output_type": "execute_result"
    }
   ],
   "source": [
    "[fuel(14),fuel(1969),fuel(100756)] == [2,966,50346]"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 47,
   "metadata": {},
   "outputs": [
    {
     "data": {
      "text/plain": [
       "4896902"
      ]
     },
     "execution_count": 47,
     "metadata": {},
     "output_type": "execute_result"
    }
   ],
   "source": [
    "modules = readdlm(\"input.txt\", '\\t', Int, '\\n')\n",
    "sum(fuel.(modules))"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 34,
   "metadata": {},
   "outputs": [
    {
     "name": "stdout",
     "output_type": "stream",
     "text": [
      "search: \u001b[0m\u001b[1mw\u001b[22m\u001b[0m\u001b[1mh\u001b[22m\u001b[0m\u001b[1mi\u001b[22m\u001b[0m\u001b[1ml\u001b[22m\u001b[0m\u001b[1me\u001b[22m\n",
      "\n"
     ]
    },
    {
     "data": {
      "text/latex": [
       "\\begin{verbatim}\n",
       "while\n",
       "\\end{verbatim}\n",
       "\\texttt{while} loops repeatedly evaluate a conditional expression, and continue evaluating the body of the while loop as long as the expression remains true. If the condition expression is false when the while loop is first reached, the body is never evaluated.\n",
       "\n",
       "\\section{Examples}\n",
       "\\begin{verbatim}\n",
       "julia> i = 1\n",
       "1\n",
       "\n",
       "julia> while i < 5\n",
       "           println(i)\n",
       "           global i += 1\n",
       "       end\n",
       "1\n",
       "2\n",
       "3\n",
       "4\n",
       "\\end{verbatim}\n"
      ],
      "text/markdown": [
       "```\n",
       "while\n",
       "```\n",
       "\n",
       "`while` loops repeatedly evaluate a conditional expression, and continue evaluating the body of the while loop as long as the expression remains true. If the condition expression is false when the while loop is first reached, the body is never evaluated.\n",
       "\n",
       "# Examples\n",
       "\n",
       "```jldoctest\n",
       "julia> i = 1\n",
       "1\n",
       "\n",
       "julia> while i < 5\n",
       "           println(i)\n",
       "           global i += 1\n",
       "       end\n",
       "1\n",
       "2\n",
       "3\n",
       "4\n",
       "```\n"
      ],
      "text/plain": [
       "\u001b[36m  while\u001b[39m\n",
       "\n",
       "  \u001b[36mwhile\u001b[39m loops repeatedly evaluate a conditional expression, and continue\n",
       "  evaluating the body of the while loop as long as the expression remains\n",
       "  true. If the condition expression is false when the while loop is first\n",
       "  reached, the body is never evaluated.\n",
       "\n",
       "\u001b[1m  Examples\u001b[22m\n",
       "\u001b[1m  ≡≡≡≡≡≡≡≡≡≡\u001b[22m\n",
       "\n",
       "\u001b[36m  julia> i = 1\u001b[39m\n",
       "\u001b[36m  1\u001b[39m\n",
       "\u001b[36m  \u001b[39m\n",
       "\u001b[36m  julia> while i < 5\u001b[39m\n",
       "\u001b[36m             println(i)\u001b[39m\n",
       "\u001b[36m             global i += 1\u001b[39m\n",
       "\u001b[36m         end\u001b[39m\n",
       "\u001b[36m  1\u001b[39m\n",
       "\u001b[36m  2\u001b[39m\n",
       "\u001b[36m  3\u001b[39m\n",
       "\u001b[36m  4\u001b[39m"
      ]
     },
     "execution_count": 34,
     "metadata": {},
     "output_type": "execute_result"
    }
   ],
   "source": [
    "?while"
   ]
  }
 ],
 "metadata": {
  "kernelspec": {
   "display_name": "Julia 1.5.2",
   "language": "julia",
   "name": "julia-1.5"
  },
  "language_info": {
   "file_extension": ".jl",
   "mimetype": "application/julia",
   "name": "julia",
   "version": "1.5.2"
  }
 },
 "nbformat": 4,
 "nbformat_minor": 4
}
