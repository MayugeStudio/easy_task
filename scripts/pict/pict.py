import os
import subprocess
import sys


def main(args: list[str]) -> None:
    if len(args) == 0:
        print("Usage: pict.py [filename]")
        sys.exit(1)
    cwd = os.getcwd()  # Project Root
    filename = os.path.normpath(args[0])
    base_filename = os.path.split(filename)[-1]
    out_ext = "csv"
    out_path_name = os.path.join(cwd, "scripts", "pict", "out", os.path.splitext(base_filename)[0] + "_out." + out_ext)
    # execute pict
    pict_out = (subprocess
                .run(["pict", filename, "/o:2"], capture_output=True, text=True)
                .stdout.replace("\t", ","))
    header = pict_out.split("\n")[0]
    body = sorted(pict_out.split("\n")[1:])
    result: list[str] = [t for t in [header] + body if len(t) != 0]

    with open(out_path_name, "w", encoding="utf-8") as f:
        f.write("\n".join(result))


if __name__ == '__main__':
    main(sys.argv[1:])
