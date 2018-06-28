# coding: utf8


import os
from PIL import Image, ImageDraw, ImageFont


def starter():

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

    main_menu()


def read_txt(file_name):

    path = str(open(file_name, "r").read())

    if len(path) > 0 and path[len(path) - 1] != "/":
        path += "/"

    return path

    main_menu()


def main_menu():

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


def logo_integrate():

    def get_logo_image(text, width_target, height_target):
        def get_image(text, font_size, color, font_path, width_target, height_target):
            image_font = ImageFont.truetype(font_path, font_size)
            height_image = sum(image_font.getmetrics())
            width_image = image_font.getsize(text)[0]
            objImage = Image.new('L', (width_image, height_image))
            ImageDraw.Draw(objImage).text((0, 0), text, fill=100, font=image_font)
            objImage = objImage.rotate(45, resample=Image.BICUBIC, expand=True)

            return objImage

        color = (255, 255, 255)
        font_path = "/home/vitjrbog/.fonts/my/impact.ttf"
        if int(width_target) > int(height_target):
            font_size = int(0.012 * int(width_target)) + 2
        else:
            font_size = int(0.020 * int(height_target)) + 2
        objLogoImage = get_image(text, font_size, color, font_path, width_target, height_target)

        return objLogoImage

    PATH = read_txt("path.txt")

    list_files_images = os.listdir(PATH + "images/")

    if len(list_files_images) < 1:
        print("\nCOMPUTER [.. Integrate a logo to images]: " +
              "Folder \"images\" is empty. " +
              "Please, copy images and try again. " +
              "Return to Main menu...")
        main_menu()

    logo = "vk.com/orb_overhear"
    i = 0
    while i < len(list_files_images):
        file_name_image = list_files_images[i]

        objTargetImage = Image.open(PATH + "images/" + file_name_image)
        (width_image, height_image) = objTargetImage.size

        if int(width_image) > int(height_image):
            width_interval = int(width_image * 0.03)
            height_interval = int(height_image * 0.12)
        else:
            width_interval = int(width_image * 0.03)
            height_interval = int(height_image * 0.12)

        objLogoImage = get_logo_image(logo, width_image, height_image)
        (width_logo, height_logo) = objLogoImage.size

        x = width_interval
        while x < int(width_image):
            y = int(height_interval / 4)
            while y < int(height_image):

                objTargetImage.paste((255, 255, 255), (x, y), objLogoImage)

                y += int(height_logo) + height_interval

            x += int(width_logo) + width_interval

        objTargetImage.save(PATH + "output/watermarked_" + file_name_image)
        print("\nCOMPUTER [.. -> Integrate a logo to images]: " +
              "Logo was successfully integrated to \"" +
              str(file_name_image) + "\".")

        i += 1

    print("\nCOMPUTER [.. -> Integrate a logo to images]: " +
          "All images was changed. Return to Main menu.")

    main_menu()


starter()
