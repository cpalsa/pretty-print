
# pretty-print.sh
A bash script that will attempt to format the user's clipboard as XML or JSON.

## Dependencies

This script relies on the following linux packages: 
 - `jq`
 - `xclip`
 - `xmlstarlet`

Using `apt` to install them looks something like this:

    sudo apt install -y jq xclip xmlstarlet
    
## Usage
The best way to use this script is with a keyboard shortcut, rather than running it directly from the command line. With that said, the rest of this usage section will assume that you have the script installed in `/usr/local/bin/pretty-print.sh` and have it bound to a keyboard shortcut of `CTRL+SHIFT+F`.

 1. Highlight the XML or JSON text you wish to format and copy it to the clipboard (`CTRL+C`)
 2. Activate the script using your keyboard shortcut (`CTRL+SHIFT+F`). You will see a notification of the script's success or failure.
 3. The clipboard contents should now be formatted! Paste the content back using `CTRL+V`.

## JSON Example

If you copied this JSON into your clipboard:

    {"widget":{"debug":"on","window":{"title":"Sample Konfabulator Widget","name":"main_window","width":500,"height":500},"image":{"src":"Images/Sun.png","name":"sun1","hOffset":250,"vOffset":250,"alignment":"center"},"text":{"data":"Click Here","size":36,"style":"bold","name":"text1","hOffset":250,"vOffset":100,"alignment":"center","onMouseUp":"sun1.opacity = (sun1.opacity / 100) * 90;"}}}

After running the script it would look like this:

    {
      "widget": {
        "debug": "on",
        "window": {
          "title": "Sample Konfabulator Widget",
          "name": "main_window",
          "width": 500,
          "height": 500
        },
        "image": {
          "src": "Images/Sun.png",
          "name": "sun1",
          "hOffset": 250,
          "vOffset": 250,
          "alignment": "center"
        },
        "text": {
          "data": "Click Here",
          "size": 36,
          "style": "bold",
          "name": "text1",
          "hOffset": 250,
          "vOffset": 100,
          "alignment": "center",
          "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
        }
      }
    }

## XML Example
If you copied this XML into your clipboard:

    <widget><debug>on</debug><window title="Sample Konfabulator Widget"><name>main_window</name><width>500</width><height>500</height></window><image src="Images/Sun.png" name="sun1"><hOffset>250</hOffset><vOffset>250</vOffset><alignment>center</alignment></image><text data="Click Here" size="36" style="bold"><name>text1</name><hOffset>250</hOffset><vOffset>100</vOffset><alignment>center</alignment><onMouseUp>sun1.opacity = (sun1.opacity / 100) * 90;</onMouseUp></text></widget>

After running the script it would look like this:

    <widget>
    	<debug>on</debug>
    	<window title="Sample Konfabulator Widget">
    		<name>main_window</name>
    		<width>500</width>
    		<height>500</height>
    	</window>
    	<image src="Images/Sun.png" name="sun1">
    		<hOffset>250</hOffset>
    		<vOffset>250</vOffset>
    		<alignment>center</alignment>
    	</image>
    	<text data="Click Here" size="36" style="bold">
    		<name>text1</name>
    		<hOffset>250</hOffset>
    		<vOffset>100</vOffset>
    		<alignment>center</alignment>
    		<onMouseUp>sun1.opacity = (sun1.opacity / 100) * 90;</onMouseUp>
    	</text>
    </widget>


