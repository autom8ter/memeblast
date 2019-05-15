package cats

import "github.com/autom8ter/smsdos/messages"

func Blast() *messages.Blast {
	return &messages.Blast{
		Messages: Messages(),
		Images:   Images(),
	}
}

func Images() messages.Images {
	return []string{
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Facidcow.com%2Fpics%2F20100607%2Fugly_cats_22.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Fimages.huffingtonpost.com%2F2014-12-08-Cat27.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.acidcow.com%2Fpics%2F20100607%2Fugly_cats_25.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fi.ytimg.com%2Fvi%2FnAvlle48MZc%2Fhqdefault.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2F2.bp.blogspot.com%2F_R7W4oG1BxAQ%2FSwzOz3ywNvI%2FAAAAAAAABK0%2FHxLw9cCmiv4%2Fs1600%2Fugly-white-cat-1.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Facidcow.com%2Fpics%2F20100607%2Fugly_cats_21.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fpics.me.me%2Fvia-9gag-com-13593726.png&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fugliestthingintheworld.com%2Fsites%2Fdefault%2Ffiles%2Fstyles%2Fthing_big%2Fpublic%2Fugliest-cat-in-the-world-sphynx-cat.jpg%3Fitok%3DQv9DoLrE&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.wallpapersafari.com%2F20%2F67%2FqSwQas.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Ftse4.mm.bing.net%2Fth%3Fid%3DOIP.cMZgZZZTfhFbsu2RdUaW6wAAAA%26pid%3DApi&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Fwww.catsaroundtheglobe.com%2Fwp-content%2Fuploads%2FCATG-one-ugly-cat-1.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Fassets.nydailynews.com%2Fpolopoly_fs%2F1.1276954.1379115970!%2Fimg%2FhttpImage%2Fimage.jpg_gen%2Fderivatives%2Fgallery_1200%2Fchinese-crested-hairless-dog.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2F3.bp.blogspot.com%2F-4AlPF1Vz44Y%2FUPPXll8-4VI%2FAAAAAAAAN98%2FIFeZHUOogGg%2Fs1600%2FFunny%2BUgly%2BDogs%2B2013_2.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fi2.wp.com%2Fearthnworld.com%2Fwp-content%2Fuploads%2F2018%2F01%2FVAMPIRE-CAT.jpg%3Fresize%3D1000%252C750%26ssl%3D1&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.wallpapersafari.com%2F20%2F83%2Fzk1awQ.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fmikesflyingpigs.files.wordpress.com%2F2011%2F04%2F2.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Ftuxedocat.us%2Fwordpress%2Fwp-content%2Fuploads%2F2015%2F03%2FNeo-laying.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.hdwallpaper.nu%2Fwp-content%2Fuploads%2F2017%2F04%2Fcat-10.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fst.depositphotos.com%2F1806346%2F2040%2Fi%2F950%2Fdepositphotos_20405629-stock-photo-cat-wearing-coat.jpg&f=1", "https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.shopify.com%2Fs%2Ffiles%2F1%2F0919%2F5220%2Fproducts%2Fcoffee2.jpg%3Fv%3D1468604975&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fthatdrop.com%2Fwp-content%2Fuploads%2F2014%2F05%2Fphoto19-1024x1024.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.catnipcamera.com%2Fwp-content%2Fuploads%2F2012%2F03%2FDSCN0537.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fimg.huffingtonpost.com%2Fasset%2F56d76a331e0000950070f2e2.jpeg%3Fcache%3Drxc3k9qny0%26ops%3D1910_1000&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fi.ytimg.com%2Fvi%2FRlZzEgmoNhw%2Fhqdefault.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fi.ytimg.com%2Fvi%2FUzD5hkQLNBk%2Fmaxresdefault.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Fwww.cultofweird.com%2Fwp-content%2Fuploads%2F2013%2F02%2Fbad-taxidermy-cat.jpg&f=1",
		"https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2F1.bp.blogspot.com%2F-LqxZe6-HPgE%2FTse_OvipS1I%2FAAAAAAAABl0%2FDKpf3Gmt38E%2Fs1600%2Fbad%2Btaxidermy_what%2Bthe%2Bcat.jpg&f=1",
	}
}

func Messages() messages.Messages {
	return []string{
		"Get cat-blasted, son",
		"Check out this cat",
		"Who let the cat out of the bag?",
		"Have you met this cat?",
		"Its raining cats",
		"All cats love fish but fear to wet their paws",
		"What's wrong, cat's got your tongue?",
		"Feel the power of this cat",
		"Cat-attack, cat, cat, cat-attack",
		"Have you been nice to a cat lately?",
	}
}
