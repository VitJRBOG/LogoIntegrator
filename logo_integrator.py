# coding: utf8


import os
from PIL import Image


def starter():
    try:
        if os.path.exists("path.txt") is False:
            file_text = open("path.txt", "w")
            file_text.write("")
            file_text.close()
            print("\nCOMPUTER: Was created file \"path.txt\".")

        PATH = read_txt("path.txt")

        if os.path.exists("logo.txt") is False:
            file_text = open("logo.txt", "w")
            file_text.write("")
            file_text.close()
            print("\nCOMPUTER: Was created file \"logo.txt\".")

        LOGO = read_txt("logo.txt")

        if os.path.exists(LOGO + "logo") is False:
            os.mkdir(str(LOGO) + "logo")
            print("\nCOMPUTER: Was created directory \"logo\".")

        if os.path.exists(PATH + "images") is False:
            os.mkdir(str(PATH) + "images")
            print("\nCOMPUTER: Was created directory \"images\".")

        if os.path.exists(PATH + "output") is False:
            os.mkdir(str(PATH) + "output")
            print("\nCOMPUTER: Was created directory \"output\".")

    except Exception as var_except:
        print("\nCOMPUTER: Error, " +
              str(var_except) +
              ". Exit from program...")
        exit()

    main_menu()


def read_txt(file_name):
    try:
        path = str(open(file_name, "r").read())

        if len(path) > 0 and path[len(path) - 1] != "/":
            path += "/"

        return path

    except Exception as var_except:
        print("\nCOMPUTER [.. -> Read " + file_name + "]: Error, " +
              str(var_except) +
              ". Return to Main menu...")

    main_menu()


def main_menu():

    try:
        print("\nCOMPUTER [Main menu]: You are in Main menu.")
        print("COMPUTER [Main menu]: 1 == Integrate a logo to images.")
        print("COMPUTER [Main menu]: 0 == Close program.")

        user_answer = raw_input("USER [Main menu]: (1/0) ")

        if user_answer == "0":
            print("COMPUTER [Main menu]: Exit from program...")
            exit()
        else:
            if user_answer == "1":
                logo_integrate()
            else:
                print("\nCOMPUTER [Main menu]: Unknown command. " +
                      "Retry query...")
                main_menu()

    except Exception as var_except:
        print("\nCOMPUTER [Main menu]: Error, " +
              str(var_except) +
              ". Exit from program...")
        exit()


def logo_integrate():

    try:
        PATH = read_txt("path.txt")
        LOGO = read_txt("logo.txt")

        list_files_images = os.listdir(PATH + "images/")

        if len(list_files_images) < 1:
            print("\nCOMPUTER [.. Integrate a logo to images]: " +
                  "Folder \"images\" is empty. " +
                  "Please, copy images and try again. " +
                  "Return to Main menu...")
            main_menu()

        i = 0
        while i < len(list_files_images):
            file_name_image = list_files_images[i]

            image = Image.open(PATH + "images/" + file_name_image)
            logo = Image.open(LOGO + "logo/logo.png")

            (width_image, height_image) = image.size
            (width_logo, height_logo) = logo.size

            x = 5
            while x < int(width_image):
                y = 5
                while y < int(height_image):

                    image.paste(logo, (x, y), logo)

                    y += int(height_logo) + 10

                x += int(width_logo) + 10

            image.save(PATH + "output/" + file_name_image)
            print("\nCOMPUTER [.. -> Integrate a logo to images]: " +
                  "Logo was successfully integrated to \"" +
                  str(file_name_image) + "\".")

            i += 1

        print("\nCOMPUTER [.. -> Integrate a logo to images]: " +
              "All images was changed. Return to Main menu.")

    except Exception as var_except:
        print("\nCOMPUTER [.. -> Integrate a logo to images]: Error, " +
              str(var_except) +
              ". Return to Main menu...")

    main_menu()


starter()
