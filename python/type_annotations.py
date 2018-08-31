from typing import Set

def my_name_in_title_case(name: str) -> str:
    return name.title()

def extract_distinct_letters_set(string: str) -> Set[str]:
    return set(string)

if __name__ == "__main__":
    tname = my_name_in_title_case("anshuman dhuliya")
    print(tname)
    tname = my_name_in_title_case("your name")
    print(f"hello {tname}")
