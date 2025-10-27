#!/usr/bin/env python3
"""Fuzz testing for critical parsers using Atheris"""

import sys
import atheris
import json
from app import process_input


@atheris.instrument_func
def fuzz_process_input(data):
    """Fuzz test the process_input function"""
    try:
        fdp = atheris.FuzzedDataProvider(data)
        
        # Generate fuzzed input
        input_string = fdp.ConsumeUnicodeNoSurrogates(fdp.ConsumeIntInRange(0, 10000))
        
        # Test the parser
        result = process_input(input_string)
        
        # Basic sanity check
        assert isinstance(result, str), "Result should be a string"
        
    except ValueError:
        # Expected exceptions for invalid input
        pass
    except Exception as e:
        # Unexpected exceptions should be reported
        print(f"Unexpected exception: {e}")
        raise


@atheris.instrument_func
def fuzz_json_parsing(data):
    """Fuzz test JSON parsing"""
    try:
        fdp = atheris.FuzzedDataProvider(data)
        
        # Generate fuzzed JSON-like data
        json_string = fdp.ConsumeUnicodeNoSurrogates(fdp.ConsumeIntInRange(0, 1000))
        
        # Try to parse as JSON
        try:
            parsed = json.loads(json_string)
        except (json.JSONDecodeError, UnicodeDecodeError):
            # Expected for invalid JSON
            pass
            
    except Exception as e:
        print(f"Unexpected exception in JSON fuzzing: {e}")
        raise


def main():
    """Run fuzz tests"""
    print("Starting fuzz testing...")
    
    # Initialize atheris
    atheris.Setup(sys.argv, fuzz_process_input)
    
    # Run fuzzing
    atheris.Fuzz()


if __name__ == '__main__':
    main()
