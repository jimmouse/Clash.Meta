package listener

func StopListener() {
	if tunLister != nil {
		_ = tunLister.Close()
		tunLister = nil
	}
	if httpListener != nil {
		_ = httpListener.Close()
		httpListener = nil
	}
	if socksListener != nil {
		_ = socksListener.Close()
		socksListener = nil
	}

	if autoRedirListener != nil {
		_ = autoRedirListener.Close()
		autoRedirListener = nil
	}

	if autoRedirListener != nil {
		_ = autoRedirListener.Close()
		autoRedirListener = nil
	}

	if tproxyListener != nil {
		_ = tproxyListener.Close()
		tproxyListener = nil
	}

	if mixedListener != nil {
		_ = mixedListener.Close()
		mixedListener = nil
	}

	if mixedListener != nil {
		_ = mixedListener.Close()
		mixedListener = nil
	}

	if shadowSocksListener != nil {
		_ = shadowSocksListener.Close()
		shadowSocksListener = nil
	}

	if shadowSocksListener != nil {
		_ = shadowSocksListener.Close()
		shadowSocksListener = nil
	}

	if vmessListener != nil {
		_ = vmessListener.Close()
		vmessListener = nil
	}

	if tuicListener != nil {
		_ = tuicListener.Close()
		tuicListener = nil
	}
}
