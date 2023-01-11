with open("bitfont.go", 'w') as gofontfile:
    bitmaps =  []
    with open("bitfont.map", 'r') as bitmapfile:
        bitmaps.extend(bitmapfile)
    gofontfile.write("package imaging\n\n") # file head
    gofontfile.write("/* This is a font file generated using goFontMaker by Paul Friedrich Vierkorn.\nFont Source File: bitfont.map\nFont Source Author: Paul Friedrich Vierkorn\nCredit String: "+bitmaps[0]+"*/\n\n")
    bitmaps.pop(0) # This removes the credits string for processing.
    gofontfile.write("var Font = map[rune][][]bool {") # start golang map
    for thisline in bitmaps:
        line = thisline.strip()
        if len(line) == 1: # single rune definition
            print("Found rune '"+line+"'!")
            gofontfile.write("\n'"+line+"' : {")
        elif len(line) == 0: # end of rune definition
            print("Ending definition for current rune.")
            gofontfile.write("},")
        else: # read bitmap, append to map
            print("Defining line for current rune: "+line)
            gofontfile.write("\n{")
            for bit in line:
                if bit == '1':
                    gofontfile.write("true,")
                if bit == '0':
                    gofontfile.write("false,")
            gofontfile.write("},")
    gofontfile.write("},\n}\n\nvar CharWidth = 8\nvar CharMiddle = 6") # supporting variables.