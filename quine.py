quine = """quine = 
quotes = '"'*3
print quine[:8] + quotes + quine + quotes
print quine[8:]"""

quotes = '"'*3
print quine[:8] + quotes + quine + quotes
print quine[8:]
