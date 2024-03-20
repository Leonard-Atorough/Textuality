using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace TinyCompiler
{
    public class Token
    {
        public string Type { get; set; }
        public string Value { get; set; }
    }

    public class Lexer
    {
        private readonly string _input;
        private int _position;
        private readonly List<Token> _tokens;

        public Lexer(string input)
        {
            _input = input;
            _tokens = new List<Token>();
        }

        public List<Token> Tokenize()
        {
            // Implement your lexer logic here to populate _tokens
            return _tokens;
        }
    }

    public class Parser
    {
        private readonly List<Token> _tokens;
        private int _position;

        public Parser(List<Token> tokens)
        {
            _tokens = tokens;
        }

        public Node Parse()
        {
            // Implement your parser logic here
            return new Node(); // Placeholder for actual AST node
        }
    }

    public class Node
    {
        // Define your AST nodes here
    }

    public class Compiler
    {
        private readonly Lexer _lexer;
        private readonly Parser _parser;

        public Compiler(string input)
        {
            _lexer = new Lexer(input);
            _parser = new Parser(_lexer.Tokenize());
        }

        public void Compile()
        {
            var ast = _parser.Parse();
            // Implement your compilation logic here
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var input = "int x = 42;";
            var compiler = new Compiler(input);
            compiler.Compile();

            Console.WriteLine("Compilation complete.");
        }
    }
}
