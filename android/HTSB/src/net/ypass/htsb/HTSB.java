package net.ypass.htsb;

import android.app.Activity;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.content.IntentFilter;
import android.os.Bundle;
import android.view.MotionEvent;
import android.view.View;
import android.view.View.OnTouchListener;
import android.widget.Button;
import android.widget.TextView;

public class HTSB extends Activity {
	private long starttime;
	private Button spacebar;
	private boolean pressed = false;
	public static final String updateTimer = "net.ypass.htsb.updateTimer";
	
	/** Called when the activity is first created. */
	@Override
	public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.main);
		spacebar = (Button)findViewById(R.id.space);
		spacebar.setOnTouchListener(onSpaceTouchEvent);
		registerReceiver(updateTimerReceiver, new IntentFilter(updateTimer));
	}
	
	private BroadcastReceiver updateTimerReceiver = new BroadcastReceiver() {
		public void onReceive(Context context, Intent intent) {
			((TextView)findViewById(R.id.time)).setText(intent.getLongExtra("curtimer", 0) + " seconds");
		}
	};
	
	private OnTouchListener onSpaceTouchEvent = new OnTouchListener() {
		public boolean onTouch(View v, MotionEvent m) {
			
			switch (m.getAction()) {
				case MotionEvent.ACTION_DOWN:
					starttime= System.currentTimeMillis()/1000;
					spacebar.setPressed(true);
					pressed = true;
					Thread t = new Thread(timer);
					t.start();
					break;
				case MotionEvent.ACTION_UP:
					endtime= System.currentTimeMillis()/1000;
					spacebar.setPressed(false);
					pressed = false;
					break;
			}
			return true;
		}
	};
	
	private Runnable timer = new Runnable() {
		public void run() {
			while (pressed == true) {
				long timer = (System.currentTimeMillis()/1000) - starttime;
				sendBroadcast(new Intent(HTSB.updateTimer).putExtra("curtimer", timer));
				try {
					this.wait(100);
				} catch (Exception e) { }
			}
		}
	};
}